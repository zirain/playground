# Coroot

https://docs.coroot.com/installation/k8s-operator

## Install Coroot operator via Helm

```shell
helm install -n coroot --create-namespace coroot-operator coroot/coroot-operator
```

```shell
kubectl port-forward -n coroot service/coroot-coroot 8080:8080  
```