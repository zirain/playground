# Install Thanos with Fluxcd

## Prerequisites

Make sure fluxcd installed via [quickstart](./quickstart.md)

## Install Thanos

helm chart repo: https://github.com/bitnami/charts/tree/main/bitnami/thanos

### Create namespace

```bash
kubectl create ns monitoring
```

### Create HelmRepository

```bash
cat <<EOF | kubectl apply -f -
apiVersion: source.toolkit.fluxcd.io/v1beta2
kind: HelmRepository
metadata:
  name: bitnami
  namespace: monitoring
spec:
  type: "oci"
  interval: 5m0s
  url: oci://registry-1.docker.io/bitnamicharts
EOF
```

### Create Thanos HelmRelease

```bash
kubectl apply -f /root/go/src/github.com/zirain/playground/k8s/fluxcd/helmreleases/thanos.yaml
```

### Create Prometheus HelmRelease

```bash
kubectl apply -f /root/go/src/github.com/zirain/playground/k8s/fluxcd/helmreleases/prometheus.yaml
```

### Check Status

```bash
kubectl describe helmrepo bitnami -n monitoring
kubectl describe hr thanos -n monitoring
kubectl describe hr prometheus -n monitoring
```


## Cleanup

```bash
kubectl delete helmrepo bitnami -n monitoring --ignore-not-found
kubectl delete hr thanos -n monitoring --ignore-not-found
kubectl delete hr prometheus -n monitoring
```