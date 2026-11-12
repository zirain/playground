#!/usr/bin/env bash
# Usage: test.sh [N] [kubeconfig dir]
#   Port protocol matrix and where N requests to `whereami` land, for each backend
#   namespace (amb, amb-wp, sc), through both EG gateways and an in-mesh client.
DEF="$(cd "$(dirname "$0")/../../istio/multicluster" && pwd)/.kube"
N=${1:-30}; K=${2:-$DEF}; KC="kubectl --kubeconfig $K/cluster1"
gw() { $KC -n envoy-gateway-system get svc -l gateway.envoyproxy.io/owning-gateway-name=$1 -o jsonpath='{.items[0].metadata.name}'; }
$KC -n envoy-gateway-system port-forward svc/$(gw ambient-endpoint) 18290:80 >/dev/null 2>&1 & P1=$!
$KC -n envoy-gateway-system port-forward svc/$(gw ambient-service) 18291:80 >/dev/null 2>&1 & P2=$!
trap 'kill $P1 $P2' EXIT; sleep 3
B=$(mktemp)
req() { local c; c=$(curl -s -m 5 -o $B -w '%{http_code}' localhost:$1$3 -H "Host: $2"); [[ $2 == nginx* ]] && c="$c $(tr -d '\n' <$B | sed 's/nginx in //')"; echo "$c"; }
mesh() { $KC -n amb exec meshclient -- curl -s -m 5 -o /dev/null -w '%{http_code}' "$1"; }
for ns in amb amb-wp sc; do
  echo "### namespace $ns"
  printf "%-14s %-26s %-26s %s\n" backend "EG Endpoint" "EG Service" "mesh client"
  for c in c1 c2; do for s in http-$c:8000:/get http2-$c:8001:/get nginx-h2c-$c:8004:/; do IFS=: read n p u <<<"$s"
    printf "%-14s %-26s %-26s %s\n" $n "$(req 18290 $n.$ns.example.com $u)" "$(req 18291 $n.$ns.example.com $u)" "$(mesh http://$n.$ns:$p$u)"
  done; done
  for g in Endpoint:18290 Service:18291; do echo "  whereami x$N via EG ${g%%:*}:"
    for i in $(seq $N); do r=$(curl -s -m 5 -w ' [%{http_code}]' localhost:${g##*:}/ -H "Host: whereami.$ns.example.com"); echo "${r//$'\n'/}"; done | sort | uniq -c | sed 's/^/   /'; done
  echo "  whereami x$N via mesh client:"
  $KC -n amb exec meshclient -- sh -c "for i in \$(seq $N); do curl -s -m 5 http://whereami.$ns:8000/ || echo fail; done" | sort | uniq -c | sed 's/^/   /'
done
