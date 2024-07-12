#!/usr/bin/env bash

set -euo pipefail

# Setup default values
ISTIO_VERSION=${ISTIO_VERSION:-$(istioctl version --remote=false)}
MIRROR_REGISTRY=${MIRROR_REGISTRY:-"192.168.3.73:5000"}

images=(pilot proxyv2 ztunnel install-cni)

for imageName in ${images[@]} ; do
    crane cp "docker.io/istio/${imageName}:${ISTIO_VERSION}" "${MIRROR_REGISTRY}/istio/${imageName}:${ISTIO_VERSION}"
    crane cp "docker.io/istio/${imageName}:${ISTIO_VERSION}-distroless" "${MIRROR_REGISTRY}/istio/${imageName}:${ISTIO_VERSION}-distroless"
done