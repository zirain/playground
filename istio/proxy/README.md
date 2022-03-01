# Build istio proxy

```shell
export ISTIO_BASE="/root/GoProjects/src/istio.io"
docker run --name envoy_build -v ${ISTIO_BASE}:/root/GoProjects/src istio-proxy-build bash -c "export BAZEL_BUILD_ARGS=--override_repository=envoy=/root/GoProjects/src/envoy;cd /root/GoProjects/src/proxy;make build_envoy"
docker cp envoy_build:/root/GoProjects/src/proxy/bazel-bin/src/envoy/envoy ${ISTIO_BASE}/istio/out/linux_amd64/release
```