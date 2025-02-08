# Telemetry: support additional labels

**Available in Istio 1.25+**

## Setup Istio

```shell
istioctl install -f iop.yaml -y
kubectl label ns default istio-injection=enabled --overwrite
kubectl apply -f apps/http
kubectl apply -f telemetry.yaml
```

## Verify

```shell
kubectl exec -it deploy/sleep -- curl httpbin:8000/get

# upstream_role and downstream_role should not be unknown
istioctl x es -oprom deploy/sleep | grep upstream_role
istioctl x es -oprom deploy/httpbin | grep downstream_role

```