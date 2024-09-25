#!/usr/bin/env bash

set -euo pipefail

# Setup default values
ISTIO_VERSION=${ISTIO_VERSION:-$(istioctl version --remote=false | cut -d ':' -f2 | xargs)}
MIRROR_REGISTRY=${MIRROR_REGISTRY:-"192.168.4.146:5000"}

istioImages=(pilot proxyv2 ztunnel install-cni)

for imageName in ${istioImages[@]} ; do
    crane cp "docker.io/istio/${imageName}:${ISTIO_VERSION}" "${MIRROR_REGISTRY}/istio/${imageName}:${ISTIO_VERSION}"
    crane cp "docker.io/istio/${imageName}:${ISTIO_VERSION}-distroless" "${MIRROR_REGISTRY}/istio/${imageName}:${ISTIO_VERSION}-distroless"
done

# MetalLB
crane cp quay.io/metallb/controller:v0.13.10 "${MIRROR_REGISTRY}/metallb/controller:v0.13.10"
crane cp quay.io/metallb/speaker:v0.13.10 "${MIRROR_REGISTRY}/metallb/speaker:v0.13.10"

# samples
crane cp kong/httpbin:latest "${MIRROR_REGISTRY}/kong/httpbin:latest"
crane cp curlimages/curl:latest "${MIRROR_REGISTRY}/curlimages/curl:latest"