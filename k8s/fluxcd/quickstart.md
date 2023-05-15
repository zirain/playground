# Fluxcd

## Install fluxcd with helm

source code: https://github.com/fluxcd-community/helm-charts

```bash
helm repo add fluxcd-community https://fluxcd-community.github.io/helm-charts
helm search repo fluxcd-community
helm show values fluxcd-community/flux2

helm install fluxcd fluxcd-community/flux2 -n fluxcd-system --create-namespace \
    -f /root/go/src/github.com/zirain/playground/k8s/fluxcd/values/fluxcd.values.yaml
```

```
kubectl logs -l app=helm-controller --tail=-1 -n fluxcd-system
kubectl logs -l app=source-controller --tail=-1 -n fluxcd-system
```

## Cleanup

```bash
helm delete fluxcd -n fluxcd-system
```