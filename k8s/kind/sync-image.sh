#!/usr/bin/env bash

set -euo pipefail

command -v regctl &>/dev/null || { echo "regctl not found. Install with: brew install regclient"; exit 1; }

ISTIO_VERSION=${ISTIO_VERSION:-$(istioctl version --remote=false | cut -d ':' -f2 | xargs)}
EG_VERSION=${EG_VERSION:-$(egctl version --remote=false | cut -d ':' -f2 | xargs)}
MIRROR_REGISTRY=${MIRROR_REGISTRY:-"192.168.4.146:5000"}
echo "ISTIO_VERSION: ${ISTIO_VERSION}, EG_VERSION: ${EG_VERSION}"

regctl registry set "${MIRROR_REGISTRY}" --tls=disabled

# sync <source-registry> <image>...  (linux/amd64 and linux/arm64 only)
sync() {
    local src=$1; shift
    for img in "$@"; do
        regctl image copy --platform linux/amd64 --platform linux/arm64 "${src}/${img}" "${MIRROR_REGISTRY}/${img}"
    done
}

# latest <registry>/<repo>: print the highest stable X.Y.Z tag (with or without a "v" prefix)
latest() {
    regctl tag ls "$1" | grep -E '^v?[0-9]+\.[0-9]+\.[0-9]+$' |
        awk '{t=$0; sub(/^v/, "", t); print t, $0}' | sort -V | tail -1 | cut -d' ' -f2
}

# Istio (distroless is used in Ambient)
for name in pilot proxyv2 ztunnel install-cni; do
    sync docker.io "istio/${name}:${ISTIO_VERSION}" "istio/${name}:${ISTIO_VERSION}-distroless"
done

sync docker.io "envoyproxy/gateway:v${EG_VERSION}"

sync docker.io \
    redis:8.6.7 \
    postgres:17 \
    golang:1.26.4 \
    kindest/node:v1.35.0 \
    kindest/node:v1.36.1 \
    kindest/node:v1.37.0 \
    kiwigrid/k8s-sidecar:1.30.9 \
    kiwigrid/k8s-sidecar:2.5.0 \
    mccutchen/go-httpbin:v2.5.0 \
    prom/prometheus:v2.52.0 \
    prom/prometheus:v3.12.0 \
    bats/bats:v1.4.1

# Always the latest stable release
for repo in \
    grafana/alloy \
    grafana/grafana \
    grafana/loki \
    grafana/tempo \
    otel/opentelemetry-collector-contrib; do
    tag=$(latest "docker.io/${repo}")
    [[ -n $tag ]] || { echo "no stable tag found for ${repo}"; exit 1; }
    sync docker.io "${repo}:${tag}"
done

sync quay.io \
    prometheus-operator/prometheus-config-reloader:v0.81.0 \
    prometheus-operator/prometheus-config-reloader:v0.89.0 \
    metallb/controller:v0.16.1 \
    metallb/speaker:v0.16.1

sync ghcr.io \
    projectcontour/yages:v0.1.0 \
    spiffe/spire-controller-manager:0.6.3 \
    spiffe/spiffe-csi-driver:0.2.7 \
    spiffe/spiffe-helper:0.11.0 \
    spiffe/oidc-discovery-provider:1.14.2 \
    spiffe/spire-agent:1.14.2

# Gateway API conformance tests
sync registry.k8s.io \
    sig-storage/csi-node-driver-registrar:v2.9.4 \
    coredns/coredns:v1.12.2 \
    gateway-api/echo-basic:v1.5.1
