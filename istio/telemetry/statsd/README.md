kubectl port-forward -nistio-system pod/`kubectl get po -nistio-system | grep statsd-influxdb-grafana | awk '{print $1}'` 8080:8888 --address 0.0.0.0

add annotations enabled metrics
```yaml
annotations:
  sidecar.istio.io/statsInclusionRegexps: .*http_local_rate_limit.*
```

```yaml
kubectl exec -it fortio-deploy-687945c6dc-98m55 -- fortio load -n 11 httpbin:8000/ip
```