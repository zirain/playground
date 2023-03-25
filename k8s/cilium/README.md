# Cilium


## Create kind cluster

```console
kind create cluster --config=/root/go/src/github.com/zirain/playground/k8s/cilium/kind-config.yaml --name cilium --image=kindest/node:v1.25.3

kind load docker-image --name cilium quay.io/cilium/cilium:v1.13.1
kind load docker-image --name cilium quay.io/cilium/operator-generic:v1.13.1
kind load docker-image --name cilium kong/httpbin
```


```console
helm repo add cilium https://helm.cilium.io/
helm repo update

helm install cilium cilium/cilium -f /root/go/src/github.com/zirain/playground/k8s/cilium/cilum-strict.yaml -nkube-system
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