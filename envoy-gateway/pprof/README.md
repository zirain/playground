


## Install Pyroscope


```shell
helm upgrade --install --create-namespace -n monitoring pyroscope grafana/pyroscope \
  -f envoy-gateway/pprof/pyroscope.yaml
```


## Enable pprof for Envoy Gateway

```shell
helm upgrade --install --create-namespace --namespace envoy-gateway-system --version v1.6.0 eg oci://docker.io/envoyproxy/gateway-helm \
  -f envoy-gateway/pprof/eg.values.yaml
```