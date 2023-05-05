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
cat <<EOF | kubectl apply -f -
apiVersion: helm.toolkit.fluxcd.io/v2beta1
kind: HelmRelease
metadata:
  name: thanos
  namespace: monitoring
spec:
  chart:
    spec:
      chart: thanos
      version: "12.5.1"
      sourceRef:
        kind: HelmRepository
        name: bitnami
  interval: 1m0s
  values:
    objstoreConfig: |-
      type: s3
      config:
        bucket: thanos
        endpoint: {{ include "thanos.minio.fullname" . }}.{{ .Release.Namespace }}.svc.cluster.local:9000
        access_key: minio
        secret_key: minio123
        insecure: true
    query:
      dnsDiscovery:
        sidecarsService: prometheus-operator-kube-p-prometheus-thanos
        sidecarsNamespace: monitoring     
    queryFrontend:
      enabled: false
    bucketweb:
      enabled: false
    compactor:
      enabled: false
    storegateway:
      enabled: true
    ruler:
      enabled: false
    metrics:
      enabled: false
    minio:
      enabled: true
      auth:
        rootPassword: minio123
        rootUser: minio
      monitoringBuckets: thanos
      accessKey:
        password: minio
      secretKey:
        password: minio123
EOF
```

### Create Prometheus HelmRelease

```bash
cat <<EOF | kubectl apply -f -
apiVersion: helm.toolkit.fluxcd.io/v2beta1
kind: HelmRelease
metadata:
  name: prometheus-operator
  namespace: monitoring
spec:
  chart:
    spec:
      chart: kube-prometheus
      version: 8.9.1
      sourceRef:
        kind: HelmRepository
        name: bitnami
  interval: 1m0s
  values:
    alertmanager:
      enabled: false
    operator:
      service:
        type: ClusterIP
    prometheus:
      disableCompaction: true
      thanos:
        create: true
        service:
          type: LoadBalancer
        objectStorageConfig:
          secretKey: objstore.yml
          secretName: thanos-objstore-secret
      service:
        type: ClusterIP
      externalLabels:
        cluster: data-producer-0
EOF
```

### Check Status

```bash
kubectl describe helmrepo bitnami -n monitoring
kubectl describe hr thanos -n monitoring
kubectl describe hr prometheus-operator -n monitoring
```


## Cleanup

```bash
kubectl delete helmrepo bitnami -n monitoring --ignore-not-found
kubectl delete hr thanos -n monitoring --ignore-not-found
kubectl delete hr prometheus-operator -n monitoring
```