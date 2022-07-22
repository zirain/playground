```console
kubectl label namespace default istio-injection=enabled --overwrite


kubectl apply --kubeconfig=/etc/karmada/karmada-apiserver.config -f networking/
```