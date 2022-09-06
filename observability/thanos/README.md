```console
kubectl create secret generic thanos-objstore --from-file=thanos.yaml=/root/thanos-config.yaml -n monitoring
kubectl create secret generic thanos-objectstorage --from-file=thanos.yaml=/root/thanos-config.yaml -n thanos
```


```console
kubectl delete secret thanos-objectstorage -nthanos
```