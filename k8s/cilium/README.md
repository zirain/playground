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