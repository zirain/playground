# Envoy gRPC Access Log Service

```shell
kubectl delete cm envoy-als -n monitoring --ignore-not-found
kubectl create cm envoy-als -n monitoring --from-file=./src
kubectl apply -f deploy.yaml -n monitoring
```