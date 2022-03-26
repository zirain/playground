# Build istio proxy

```shell
sh ./build-image.sh

export BAZEL_BUILD_ARGS=--override_repository=envoy=/root/GoProjects/src/envoy;cd /root/GoProjects/src/proxy;make build_envoy

docker cp envoy_build:/root/GoProjects/src/proxy/bazel-bin/src/envoy/envoy ${ISTIO_BASE}/istio/out/linux_amd64/release
```