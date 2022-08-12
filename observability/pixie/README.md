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
time="2022-08-11T06:39:40Z" level=info msg="Please go to 'https://work.dev.withpixie.dev/oauth/kratos/self-service/recovery/methods/link?flow=85dd96b8-04b2-4e35-bd39-5be980315269&token=dZcDLGrYb8fnc66bgbVrKxxiocGOOKDM' to set password for 'admin@default.com'"

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
helm install pixie pixie-operator/pixie-operator-chart --set deployKey="px-dep-31a4d0e4-56d8-4372-86c7-02cf44d2adbc" --set clusterName=default-cluster --namespace pl --create-namespace --set devCloudNamespace=plc
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