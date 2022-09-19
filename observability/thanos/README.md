```console
kubectl create secret generic thanos-objstore --from-file=thanos.yaml=/root/thanos-config.yaml -n monitoring
kubectl create secret generic thanos-objectstorage --from-file=thanos.yaml=/root/thanos-config.yaml -n thanos
```


```console
kubectl delete secret thanos-objectstorage -nthanos
```


```
helm install grafana grafana/grafana -ngrafana  --set service.type=LoadBalancer --set admin.password=admin
```


kubectl port-forward -nmonitoring pod/prometheus-thanos-0 1313:9090 --address 0.0.0.0
