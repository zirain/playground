# Retina

[Retina](https://retina.sh/) is a cloud-agnostic, open-source Kubernetes Network Observability platform which enables the use of Hubble as a control plane regardless of the underlying OS or CNI.


## Get latest version

```shell
curl -sL https://api.github.com/repos/microsoft/retina/releases/latest | jq -r .name
```

## Install Retina

```shell
helm upgrade --install retina oci://ghcr.io/microsoft/retina/charts/retina \
    --version v0.0.32 \
    --namespace kube-system \
    -f retina.values.yaml
```

## Install Retina Hubble

**Remove the previous release if it exists**

```shell
helm upgrade --install retina oci://ghcr.io/microsoft/retina/charts/retina-hubble \
    --version v0.0.32 \
    --namespace kube-system \
    -f retina-hubble.values.yaml
```