# Contour


## Quick start

Install contour:

```console
kubectl apply -f https://projectcontour.io/quickstart/contour.yaml
```

Install Contour Gateway Provisioner:

```console
kubectl apply -f https://projectcontour.io/quickstart/contour-gateway-provisioner.yaml
```

Create Gateway:

```console
kubectl apply -f k8s/contour-ingress/contour.yaml -nprojectcontour
```

## Verify

Create HTTPRoute:

```console
kubectl apply -f k8s/contour-ingress/httpbin-gateway.yaml
```

Verify traffic:

```
curl $(kubectl get svc envoy-contour-gateway -nprojectcontour -o jsonpath="{.status.loadBalancer.ingress[0].ip}"):8080/get
```