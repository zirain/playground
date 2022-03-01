#!/bin/bash

#docker build -t istio-proxy-build:latest .


ISTIO_BASE="/root/GoProjects/src/istio.io"
ISTIO_BUILD_TOOL=gcr.io/istio-testing/build-tools-proxy:master-2022-02-25T16-33-58

docker rm -f envoy_build
docker run --name envoy_build -v ${ISTIO_BASE}:/root/GoProjects/src ${ISTIO_BUILD_TOOL} bash -c "export BAZEL_BUILD_ARGS=--override_repository=envoy=/root/GoProjects/src/envoy;cd /root/GoProjects/src/proxy;make build_envoy"
docker cp envoy_build:/root/GoProjects/src/proxy/bazel-bin/src/envoy/envoy ${ISTIO_BASE}/istio/out/linux_amd64/release