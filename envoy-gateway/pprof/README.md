


## Install Pyroscope


```shell
helm upgrade --install --create-namespace -n monitoring pyroscope grafana/pyroscope \
  -f envoy-gateway/pprof/pyroscope.yaml
```


## Enable pprof for Envoy Gateway

```shell
export EG_VERSION="v1.6.0" # use v0.0.0-latest for latest main build
helm upgrade --install --create-namespace --namespace envoy-gateway-system --version $EG_VERSION eg oci://docker.io/envoyproxy/gateway-helm \
  -f envoy-gateway/pprof/eg.values.yaml
```