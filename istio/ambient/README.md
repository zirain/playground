# Istio Ambient Mode

```shell
istioctl -f iop.yaml -y

kubectl apply -f ../apps/http

kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.2.1/standard-install.yaml

```

```shell
kubectl label namespace default istio.io/dataplane-mode=ambient
```


## downstream -> waypoint -> upstream

1. sleep -- hbone -> waypoint  --> httpbin mesh

1. sleep -- tcp -> waypoint --> httpbin.not-ambient(not-mesh)

1. sleep(not mesh) --> waypoint --> httpbin mesh ?

1. ingress -> waypoint --> httpbin.not-ambient(not-ambient)