# OpenTelemetry


## How to build otelcontribcol


```console
COMPONENT=otelcontribcol make docker-component
```


## OpenTelemetry Demo

```console
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm install my-otel-demo open-telemetry/opentelemetry-demo
```

more details: https://github.com/open-telemetry/opentelemetry-demo/blob/v1.0.0/docs/kubernetes_deployment.md