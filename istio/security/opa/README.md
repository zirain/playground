# Istio with OPA

https://istio.io/latest/blog/2024/l7-policy-with-opa/

## Setup OPA and Istio

```shell
istioctl install -f istio/security/opa/iop.yaml
```

```shell
kubectl create ns sidecar
kubectl label ns sidecar istio-injection=enabled --overwrite
kubectl apply -f istio/security/opa/httpbin.yaml -n sidecar

kubectl apply -f istio/security/opa/policy.yaml
```

## Simple Rule

```shell
kubectl apply -f istio/security/opa/opa-cm1.yaml
kubectl rollout restart deployment -n opa
```

```shell
kubectl exec -it -n sidecar deployments/sleep -- curl -w "\nhttp_code=%{http_code}" httpbin:8000/get
kubectl exec -it -n sidecar deployments/sleep -- curl -s -w "\nhttp_code=%{http_code}" httpbin:8000/get -H "x-force-authorized: enabled"
```

## Advanced Rule

```shell
kubectl apply -f istio/security/opa/opa-cm2.yaml
kubectl rollout restart deployment -n opa
```

```shell
kubectl exec -it -n sidecar deployments/sleep -- curl -w "\nhttp_code=%{http_code}" httpbin:8000/get
kubectl exec -it -n sidecar deployments/sleep -- curl -w "\nhttp_code=%{http_code}" httpbin:8000/get -H "x-force-unauthenticated: enabled"
kubectl exec -it -n sidecar deployments/sleep -- curl -s -w "\nhttp_code=%{http_code}" httpbin:8000/get -H "x-force-authorized: true"
```