
```
kubectl apply -f /root/go/src/playground/observability/grafana-elb.yaml -nistio-system


kubectl port-forward -nistio-system pod/`kubectl get po -nistio-system | grep zipkin | awk '{print $1}'` 1313:9411  --address 0.0.0.0
kubectl port-forward -nistio-system pod/`kubectl get po -nistio-system | grep jaeger | awk '{print $1}'` 1313:16686  --address 0.0.0.0
```
