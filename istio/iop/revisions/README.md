

```console
kubectl create ns blue
kubectl create ns green


kubectl label ns blue istio-discovery=enabled istio.io/rev=blue
kubectl label ns green istio-discovery=enabled istio.io/rev=green
```