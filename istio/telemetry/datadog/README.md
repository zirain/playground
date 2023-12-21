# Istio with Datadog

## Install Istio

```shell
istioctl install -f istio/telemetry/datadog/iop.yaml -y

kubectl apply -f istio/telemetry/datadog/telemetry.yaml
```

## Install Datadog(otel-collector-contrib)

```shell
helm upgrade --install otel-collector open-telemetry/opentelemetry-collector -f istio/telemetry/datadog/values.yaml -n istio-system --create-namespace
```
