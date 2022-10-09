
```shell

helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update

helm template istio-ingress istio/gateway --namespace istio-ingress --post-renderer ./kustomize.sh

```



