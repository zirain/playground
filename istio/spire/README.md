# Spire with Istio

## Install Spire

```shell
helm upgrade --install -n spire-server spire-crds spire-crds --repo https://spiffe.github.io/helm-charts-hardened/ --create-namespace
helm upgrade --install -n spire-server spire spire --repo https://spiffe.github.io/helm-charts-hardened/ --wait --set global.spire.trustDomain="example.org"
```

## Install Istio

```shell
istioctl install -f istio/spire/iop.yaml -y
```

## Install curl and httpbin

```shell
kubectl apply -f istio/spire/curl.yaml -n spire-sample
kubectl apply -f istio/spire/httpbin.yaml -n spire-sample
```
