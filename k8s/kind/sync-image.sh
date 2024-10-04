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

golangImages=(golang:1.23.1 golang:1.22.6)
for imageName in ${golangImages[@]} ; do
    # https://github.com/google/go-containerregistry/issues/2016
    crane index filter "${imageName}" --platform linux/amd64 --platform linux/arm64 -t "${MIRROR_REGISTRY}/${imageName}"
done

# sync image from docker.io
images=(fluent/fluent-bit:2.1.4 mccutchen/go-httpbin:v2.5.0 kong/httpbin:latest curlimages/curl:latest grafana/grafana:11.0.0 prom/prometheus:v2.52.0 grafana/tempo:2.1.1 bats/bats:v1.4.1)
for imageName in ${images[@]} ; do
    crane cp "${imageName}" "${MIRROR_REGISTRY}/${imageName}"
done

# sync images from quary.io
queryImages=(prometheus-operator/prometheus-config-reloader:v0.73.2 metallb/controller:v0.13.10 metallb/speaker:v0.13.10)
for imageName in ${queryImages[@]} ; do
    crane cp "quay.io/${imageName}" "${MIRROR_REGISTRY}/${imageName}"
done
