# Fluent-bit on Kubernetes

## Setup Fluentd-bits

```console
helm repo add fluent https://fluent.github.io/helm-charts
```

```console
helm upgrade --install fluent-bit fluent/fluent-bit -f k8s/fluent-bit/values.yaml -n monitoring --create-namespace
```
## Setup Loki

```
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/addons/loki.yaml -n monitoring
```

## Setup Grafana

```console
kubectl port-forward svc/grafana -nistio-system 3000:3000 --address 0.0.0.0
```

```console
kubectl delete -f https://raw.githubusercontent.com/istio/istio/master/samples/addons/loki.yaml -n monitoring
helm uninstall fluent-bit -n monitoring
```