# Install Istio with Helm

```shell

helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update


helm install istio-base istio/base -n istio-system

#helm install istiod istio/istiod -n istio-system --wait

helm install istiod istio/istiod -n istio-system --wait -f /root/go/src/github.com/zirain/playground/istio/helm/meshconfig.yaml

```