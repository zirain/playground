# OpenTelemetry Operator

https://github.com/open-telemetry/opentelemetry-operator



## Install

```console

kubectl apply -f https://github.com/open-telemetry/opentelemetry-operator/releases/latest/download/opentelemetry-operator.yaml

kubectl apply -f /root/go/src/github.com/zirain/playground/observability/opentelemetry-operator/agent.yaml
```

## Cleanup

```console
kubectl delete -f /root/go/src/github.com/zirain/playground/observability/opentelemetry-operator/agent.yaml
kubectl delete -f https://github.com/open-telemetry/opentelemetry-operator/releases/latest/download/opentelemetry-operator.yaml
```