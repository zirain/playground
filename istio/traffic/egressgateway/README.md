# Istio with Egress Gateway

```shell
istioctl install -f istio/traffic/egressgateway/iop.yaml -y
```


```shell
kubectl apply -f istio/traffic/egressgateway/egress-http.yaml
```

```shell
kubectl apply -f istio/traffic/egressgateway/egress-authz.yaml
```