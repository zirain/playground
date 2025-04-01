# Istio Ambient Mode

```shell
istioctl -f iop.yaml -y

kubectl apply -f apps/http

kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.2.1/standard-install.yaml

```

```shell
kubectl label namespace default istio.io/dataplane-mode=ambient
```
