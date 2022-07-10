# Gateway API


## Install CRD

```console
kubectl get crd gateways.gateway.networking.k8s.io || { kubectl kustomize "github.com/kubernetes-sigs/gateway-api/config/crd?ref=v0.4.3" | kubectl apply -f -; }
```


## Setup Httpbin

1. deploy httpbin app:

```console
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.14/samples/httpbin/httpbin.yaml
```

2. deploy the Gateway API configuration:

```console
kubectl apply -f istio/traffic/gatewayapi/httpbin-gateway.yaml
```


```
export INGRESS_HOST=$(kubectl get gateways.gateway.networking.k8s.io httpbin-gateway -n istio-system -ojsonpath='{.status.addresses[*].value}')

# 200 OK
curl -s -I -HHost:httpbin.example.com "http://$INGRESS_HOST/get"

# 404 OK
curl -s -I -HHost:httpbin.example.com "http://$INGRESS_HOST/headers"
```


```console
istioctl pc log $(kubectl get po -l app=httpbin | grep httpbin | awk '{print $1}') --level debug
curl -s -I -HHost:httpbin.example.com "http://$INGRESS_HOST/get"
kubectl logs -l app=httpbin -cistio-proxy --tail=-1 | grep "my-added-header"

2022-07-10T09:41:06.417837Z	trace	envoy http	[C77179] completed header: key=my-added-header value=added-value
```