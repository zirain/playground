# Telemetry: support additional labels

**Available in Istio 1.25+**

## Setup Istio

```shell
istioctl install -f iop.yaml -y
kubectl label ns default istio-injection=enabled --overwrite

kubectl apply -f telemetry.yaml
```

## Verify HTTP

```shell
kubectl apply -f apps/http
kubectl exec -it deploy/sleep -- curl httpbin:8000/get

# upstream_role and downstream_role should not be unknown
istioctl x es -oprom deploy/sleep | grep upstream_role | head -n 1
istioctl x es -oprom deploy/httpbin | grep downstream_role | head -n 1

```


## Verify TCP

```shell
kubectl apply -f apps/tcp
kubectl exec -it deployments/busybox -- sh -c "echo world | nc tcp-echo 9000"

# upstream_role and downstream_role should not be unknown
istioctl x es -oprom deploy/busybox | grep upstream_role | head -n 1
istioctl x es -oprom deploy/tcp-echo | grep downstream_role | head -n 1

```