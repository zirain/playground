# Install OTel-collector with k8s attribute processor


```shell
helm upgrade --install otel-collector open-telemetry/opentelemetry-collector -f envoy-gateway/otel-collector/helm-values.yaml -n monitoring --create-namespace --version 0.117.3
```