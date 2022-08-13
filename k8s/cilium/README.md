# Cilium


## Create kind cluster

```console
kind create cluster --config=kind-config.yaml --name cilium --image=kindest/node:v1.23.4
```


```console
helm install cilium cilium/cilium -f cilium.yaml -nkube-system
helm upgrade cilium cilium/cilium -f cilium.yaml -nkube-system

cilium status
```

## Install demo app

```console
kubectl create -f https://raw.githubusercontent.com/cilium/cilium/v1.12/examples/minikube/http-sw-app.yaml
```


## Clean

```
helm uninstall cilium
```


## Summary

- Hubble需要基于Cilium CNI，无法像Pixie一样单独部署
- Hubble UI默认不支持多集群
- 目前Tetragon与Hubble没有继承