# Install Istio with Helm

```shell

helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update

# 安装istio-1.13.0
helm install istio-base istio/base -n istio-system --version 1.13.0
helm install istiod istio/istiod -n istio-system --wait --version 1.13.0 -f /root/go/src/github.com/zirain/playground/istio/helm/meshconfig.yaml

# 升级istio-1.17.2
helm upgrade istiod istio/istiod -n istio-system --wait --version 1.17.2 -f /root/go/src/github.com/zirain/playground/istio/helm/meshconfig.yaml

kubectl rollout restart deploy/httpbin
```