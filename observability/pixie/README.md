# Pixie

Pixie is an open source observability tool for Kubernetes applications. Pixie uses eBPF to automatically capture telemetry data without the need for manual instrumentation.

# Quick started

*** Pixie does not support kind ***

https://docs.pixielabs.ai/installing-pixie/requirements/#kubernetes-local-development-environments


## Setup cluster with k0s

https://docs.px.dev/installing-pixie/setting-up-k8s/k0s-setup/


```console
# Download k0s
curl -sSLf https://get.k0s.sh | sudo sh

# Install k0s as a service
sudo k0s install controller --single

# Start k0s as a service
sudo k0s start

# Check service, logs and k0s status
sudo k0s status

# Export Admin's Kubeconfig file
sudo k0s kubeconfig admin > /root/.kube/config
```

### Setup local pv with OpenEBS

Self-Hosted Pixie needs a PersistentVolume, we use OpenESB:

```console
kubectl apply -f https://openebs.github.io/charts/openebs-operator.yaml
```

```console
cat <<EOF | kubectl apply -f -
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: local-hostpath
  annotations:
    storageclass.kubernetes.io/is-default-class: 'true'
    openebs.io/cas-type: local
    cas.openebs.io/config: |
      - name: StorageType
        value: hostpath
      - name: BasePath
        value: /var/local-hostpath
provisioner: openebs.io/local
reclaimPolicy: Delete
volumeBindingMode: WaitForFirstConsumer
EOF
```

## Install Self-Hosted Pixie

https://docs.px.dev/installing-pixie/install-guides/self-hosted-pixie#1.-deploy-pixie-cloud

### Install Pixie cloud

```console
kubectl create namespace plc
```

```console
cd ~/go/src/pixie-io/pixie && ./dev_dns_updater --domain-name="dev.withpixie.dev"  --kubeconfig=$HOME/.kube/config --n=plc`
```

Deploy Pixie Cloud dependencies(postgres, es, nats etc.):
```
kustomize build k8s/cloud_deps/base/elastic/operator | kubectl apply -f -
kustomize build k8s/cloud_deps/public | kubectl apply -f -
```

Deploy Pixie Cloud:
```
kustomize build k8s/cloud/public/ | kubectl apply -f -
```

```
k logs create-admin-job-42gkp -nplc
```
time="2022-08-17T07:46:45Z" level=info msg="Please go to 'https://work.dev.withpixie.dev/oauth/kratos/self-service/recovery/methods/link?flow=8faf0adc-740f-4167-8a91-1a124ded449c&token=iSeO6VmjOOBgUs2ks1J9n2mF6fYWuFAb' to set password for 'admin@default.com'"

Visit `https://work.dev.withpixie.dev/auth/login?local_mode=true` get access token:

```console
px auth login --manual
```

run following command:
```
helm repo add pixie-operator https://pixie-operator-charts.storage.googleapis.com
# Get latest information about Pixie chart.
helm repo update
# Install Pixie with a memory limit for the PEM pods (per node). 2Gi is the default, 1Gi is the minimum recommended.
helm install pixie pixie-operator/pixie-operator-chart --set deployKey="<px-dep-*>" --set clusterName=default-cluster --namespace pl --create-namespace --set devCloudNamespace=plc
```

## Cleanup


```console
kustomize build k8s/cloud/public/ | kubectl delete -f -
kustomize build k8s/cloud_deps/base/elastic/operator | kubectl delete -f -
kustomize build k8s/cloud_deps/public | kubectl delete -f -
```

```console
sudo k0s stop
sudo k0s reset
```

## DNS

```
# pixie
159.138.129.168   dev.withpixie.dev work.dev.withpixie.dev segment.dev.withpixie.dev docs.dev.withpixie.dev
159.138.129.168   cloud.dev.withpixie.dev
```