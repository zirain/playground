#!/usr/bin/env bash
#
# Run on the machine where the registry container runs.
# Counts how often each tag was pulled in the last SINCE (from `docker logs`),
# and with DELETE=true deletes the tags nobody pulled, then garbage-collects.
#
#   ./mirror-cleanup.sh                  # report only
#   DELETE=true ./mirror-cleanup.sh      # delete unused tags + garbage-collect
#
# Env:
#   CONTAINER  registry container name/ID (default: the running container whose image is registry*)
#   REGISTRY   registry address from this host (default: localhost:5000)
#   SINCE      log window (default: 720h = 30 days)
#
# Counted: successful GET/HEAD manifest requests by tag, except those from regctl
# (sync-image.sh). Tags pushed within the window are kept even with 0 pulls.
# Needs: docker, curl, awk.

# scp -O k8s/kind/mirror-cleanup.sh <user>@192.168.4.146:~/

set -euo pipefail

REGISTRY=${REGISTRY:-"localhost:5000"}
SINCE=${SINCE:-720h}
DELETE=${DELETE:-false}
CONTAINER=${CONTAINER:-$(docker ps --format '{{.ID}} {{.Image}}' | awk '$2 ~ /(^|\/)registry(:|$)/ {print $1; exit}')}
[[ -n $CONTAINER ]] || { echo "registry container not found; set CONTAINER"; exit 1; }

CONFIG=/etc/distribution/config.yml # registry:3
docker exec "$CONTAINER" test -f "$CONFIG" || CONFIG=/etc/docker/registry/config.yml # registry:2

ACCEPT=(
    -H "Accept: application/vnd.oci.image.index.v1+json"
    -H "Accept: application/vnd.oci.image.manifest.v1+json"
    -H "Accept: application/vnd.docker.distribution.manifest.list.v2+json"
    -H "Accept: application/vnd.docker.distribution.manifest.v2+json"
)

# JSON string array field -> one value per line, e.g. json_list repositories
# (prints nothing for null/missing, e.g. a repo whose tags were all deleted)
json_list() {
    { grep -o "\"$1\":\[[^]]*\]" || true; } | sed -e 's/^[^[]*\[//' -e 's/\]$//' | tr ',' '\n' | tr -d '"' | sed '/^$/d'
}

digest_of() { # <repo> <tag>
    curl -fsSI "${ACCEPT[@]}" "http://${REGISTRY}/v2/$1/manifests/$2" |
        awk -F': ' 'tolower($1) == "docker-content-digest" {print $2}' | tr -d '\r'
}

tmp=$(mktemp -d); trap 'rm -rf "$tmp"' EXIT

echo "container: ${CONTAINER} (created $(docker inspect -f '{{.Created}}' "$CONTAINER")), window: ${SINCE}"
echo "note: logs before the container was created are gone, so a younger container under-counts pulls"
echo

docker logs --since "$SINCE" "$CONTAINER" > "$tmp/logs" 2>&1

# the catalog is paginated (max 1000 per page)
last=""
while :; do
    page=$(curl -fsS "http://${REGISTRY}/v2/_catalog?n=1000${last:+&last=${last}}" | json_list repositories)
    [[ -n $page ]] || break
    echo "$page"
    last=$(tail -1 <<< "$page")
    (( $(wc -l <<< "$page") < 1000 )) && break
done > "$tmp/repos"

while read -r repo; do
    curl -fsS "http://${REGISTRY}/v2/${repo}/tags/list" | json_list tags | sed "s|^|${repo}:|"
done < "$tmp/repos" > "$tmp/tags"
[[ -s $tmp/tags ]] || { echo "no tags found in ${REGISTRY}"; exit 1; }

# "<pulls> <pushed-in-window> <repo:tag>" per tag. The registry logs each request
# twice (logrus "response completed" + access log); access lines are only used
# when there are no logrus lines.
awk '
function add(method, uri, status, ua,    ref) {
    if (status !~ /^2/) return
    if (!match(uri, /^\/v2\/.+\/manifests\/[^\/?]+/)) return
    sub(/^\/v2\//, "", uri); sub(/\?.*/, "", uri)
    ref = uri; sub(/\/manifests\//, ":", ref)
    if (ref ~ /:sha256:/) return
    if ((method == "GET" || method == "HEAD") && ua !~ /regclient/) pulls[src, ref]++
    if (method == "PUT") pushed[src, ref] = 1
}
function field(re,    s) {
    if (!match($0, re)) return ""
    s = substr($0, RSTART, RLENGTH); sub(/^[^=]*="?/, "", s); sub(/"$/, "", s)
    return s
}
FNR == NR { all[$0] = 1; next }
/msg="response completed"/ {
    src = "logrus"; seen["logrus"] = 1
    add(field("http\\.request\\.method=[A-Z]+"), field("http\\.request\\.uri=\"[^\"]*\""),
        field("http\\.response\\.status=[0-9]+"), field("http\\.request\\.useragent=\"[^\"]*\""))
    next
}
match($0, /"(GET|HEAD|PUT) \/v2\/[^ ]+ HTTP\/[0-9.]+" [0-9]+/) {
    src = "access"
    split(substr($0, RSTART + 1, RLENGTH - 1), r, /[ "]+/)
    n = split($0, q, "\"")
    add(r[1], r[2], r[4], q[n - 1])
}
END {
    src = ("logrus" in seen) ? "logrus" : "access"
    for (ref in all) print pulls[src, ref] + 0, pushed[src, ref] + 0, ref
}' "$tmp/tags" "$tmp/logs" | sort -k1,1n -k3 > "$tmp/report"

# Deleting is by digest, which removes every tag pointing at it, so a digest
# shared with a tag we keep is never deleted.
while read -r count pushed ref; do
    (( count == 0 && ! pushed )) && continue
    digest_of "${ref%:*}" "${ref##*:}"
done < "$tmp/report" | sort -u > "$tmp/keep-digests"

printf '%8s  %s\n' PULLS IMAGE
while read -r count pushed ref; do
    note=""
    if (( count == 0 )); then
        if (( pushed )); then
            note="  (pushed recently, kept)"
        else
            repo=${ref%:*}
            digest=$(digest_of "$repo" "${ref##*:}")
            if grep -qxF "$digest" "$tmp/keep-digests"; then
                note="  (unused, shares image with a used tag, kept)"
            elif [[ $DELETE == true ]]; then
                curl -fsS -X DELETE "http://${REGISTRY}/v2/${repo}/manifests/${digest}" -o /dev/null
                note="  (deleted)"
            else
                note="  (unused)"
            fi
        fi
    fi
    printf '%8d  %s%s\n' "$count" "$ref" "$note"
done < "$tmp/report"

if [[ $DELETE == true ]]; then
    echo
    echo "garbage-collecting (${CONFIG})..."
    docker exec "$CONTAINER" registry garbage-collect --delete-untagged "$CONFIG"
fi
