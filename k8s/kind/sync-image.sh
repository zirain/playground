#!/usr/bin/env bash

set -euo pipefail

# Setup default values
ISTIO_VERSION=${ISTIO_VERSION:-$(istioctl version --remote=false | cut -d ':' -f2 | xargs)}
MIRROR_REGISTRY=${MIRROR_REGISTRY:-"192.168.4.146:5000"}

# sync Istio images(distroless is used in Ambient)
istioImages=(pilot proxyv2 ztunnel install-cni)
for imageName in ${istioImages[@]} ; do
    crane cp "docker.io/istio/${imageName}:${ISTIO_VERSION}" "${MIRROR_REGISTRY}/istio/${imageName}:${ISTIO_VERSION}"
    crane cp "docker.io/istio/${imageName}:${ISTIO_VERSION}-distroless" "${MIRROR_REGISTRY}/istio/${imageName}:${ISTIO_VERSION}-distroless"
done

bookinfoTag=${BOOKINFO_IMAGE_TAG:-""}
bookinfoImages=(istio/examples-bookinfo-productpage-v1
                 istio/examples-bookinfo-details-v1
                 istio/examples-bookinfo-ratings-v1
                 istio/examples-bookinfo-reviews-v1
                 istio/examples-bookinfo-reviews-v2
                 istio/examples-bookinfo-reviews-v3)
for imageName in ${bookinfoImages[@]} ; do
    crane cp "docker.io/${imageName}" "${MIRROR_REGISTRY}/${imageName}:${bookinfoTag}"
done

golangImages=(golang:1.23.1 golang:1.22.6 curlimages/curl:latest)
for imageName in ${golangImages[@]} ; do
    # https://github.com/google/go-containerregistry/issues/2016
    crane index filter "${imageName}" --platform linux/amd64 --platform linux/arm64 -t "${MIRROR_REGISTRY}/${imageName}"
done

# sync image from docker.io
images=(redis:6.0.6
        kindest/node:v1.33.1
        otel/opentelemetry-collector-contrib:0.121.0
        mccutchen/go-httpbin:v2.5.0
        kong/httpbin:latest
        grafana/alloy:v1.4.3
        grafana/grafana:11.0.0
        prom/prometheus:v2.52.0
        prom/prometheus:v3.5.0
        grafana/tempo:2.1.1
        bats/bats:v1.4.1)
for imageName in ${images[@]} ; do
    crane cp "${imageName}" "${MIRROR_REGISTRY}/${imageName}"
done

# sync images from quary.io
queryImages=(prometheus-operator/prometheus-config-reloader:v0.73.2
             metallb/controller:v0.13.10
             metallb/speaker:v0.13.10)
for imageName in ${queryImages[@]} ; do
    crane cp "quay.io/${imageName}" "${MIRROR_REGISTRY}/${imageName}"
done

# sync image from ghcr.io
ghcrImages=(projectcontour/yages:v0.1.0)
for imageName in ${ghcrImages[@]} ; do
    crane cp "ghcr.io/${imageName}" "${MIRROR_REGISTRY}/${imageName}"
done

# sync Envoy Gateway images
EG_VERSION=${EG_VERSION:-$(egctl version --remote=false | cut -d ':' -f2 | xargs)}
echo "EG_VERSION: ${EG_VERSION}"

egImages=(envoyproxy/gateway)
for imageName in ${egImages[@]} ; do
    crane cp "docker.io/${imageName}:v${EG_VERSION}" "${MIRROR_REGISTRY}/${imageName}:${EG_VERSION}"
done

# images for Gateway API conformance tests
image=(
    gcr.io/k8s-staging-gateway-api/echo-basic:v20240412-v1.0.0-394-g40c666fd
    registry.k8s.io/coredns/coredns:v1.12.2
)
for img in ${image[@]}; do
    crane cp "${img}" "${MIRROR_REGISTRY}/${img}"
done