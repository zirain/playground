# Tetragon

Cilium’s new Tetragon component enables powerful realtime, eBPF-based Security Observability and Runtime Enforcement.

## Creat cluster


## Install Tetragon via helm

```console
helm repo add cilium https://helm.cilium.io
helm repo update

helm install tetragon cilium/tetragon -n kube-system
kubectl rollout status -n kube-system ds/tetragon -w
```

## Deploy Demo Application


```console
kubectl create -f https://raw.githubusercontent.com/cilium/cilium/v1.11/examples/minikube/http-sw-app.yaml
```

To start printing events run:
```
kubectl logs -n kube-system ds/tetragon -c export-stdout -f | tetragon observe --namespace default --pod xwing
```