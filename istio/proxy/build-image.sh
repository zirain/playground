#!/bin/bash

#docker build -t istio-proxy-build:latest .


ISTIO_BASE="/root/GoProjects/src/istio.io"
CACHE_PATH="/tmp/envoy_build/cache"
RELEASE_PATH="/tmp/envoy_build/release"
ISTIO_BUILD_TOOL=gcr.io/istio-testing/build-tools-proxy:master-2022-02-25T16-33-58

mkdir -p ${CACHE_PATH}
mkdir -p ${RELEASE_PATH}
docker rm -f envoy_build
docker run --name envoy_build \
    -v ${CACHE_PATH}:/root/.cache \
    -v ${ISTIO_BASE}:/root/GoProjects/src \
    ${ISTIO_BUILD_TOOL} bash -c "sleep infinity"
    #${ISTIO_BUILD_TOOL} bash -c "export BAZEL_BUILD_ARGS=--override_repository=envoy=/root/GoProjects/src/envoy;cd /root/GoProjects/src/proxy;make build_envoy"
# copy 
# docker cp envoy_build:/root/GoProjects/src/proxy/bazel-bin/src/envoy/envoy ${RELEASE_PATH}
# docker cp envoy_build:/root/.cache ${CACHE_PATH}