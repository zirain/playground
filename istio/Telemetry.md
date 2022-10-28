# Istio Telemetry

## grafana

```console
make gen-addons && kubectl apply -f samples/addons/grafana.yaml && kubectl rollout restart deployment grafana -nistio-system

# port-forward
kubectl port-forward -nistio-system pod/`kubectl get po -nistio-system | grep grafana | awk '{print $1}'` 3000:3000 --address 0.0.0.0

# elb
kubectl apply -f /root/go/src/playground/observability/grafana-elb.yaml -nistio-system
```


## zipkin/jaeger

```
kubectl port-forward -nistio-system pod/`kubectl get po -nistio-system | grep zipkin | awk '{print $1}'` 1313:9411  --address 0.0.0.0
kubectl port-forward -nistio-system pod/`kubectl get po -nistio-system | grep jaeger | awk '{print $1}'` 1313:16686  --address 0.0.0.0
```