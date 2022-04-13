# install Loki

```console
kubectl create ns loki


helm upgrade --install loki --namespace=loki grafana/loki
```


# install Tempo

```console
kubectl create ns tempo

helm upgrade --install tempo grafana/tempo -n tempo
helm upgrade -f single-binary-grafana-values.yaml --install grafana grafana/grafana
kubectl create -f single-binary-extras.yaml
```


```

kubectl apply -f istio/telemetry/grafana/otel.yaml 
kubectl apply -f istio/telemetry/grafana/telemetry.yaml 
kubectl delete po  -nistio-system -l app=opentelemetry,component=otel-collector

```