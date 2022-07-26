# How to dev karmada-controller-manager


```console
# build karmada-controller-manager
VERSION='latest' make image-karmada-controller-manager

# load image to kind 
kind load docker-image swr.ap-southeast-1.myhuaweicloud.com/karmada/karmada-controller-manager:latest --name karmada-host

# rollout karmada-controller-manager 
kubectl rollout restart deployment/karmada-controller-manager -nkarmada-system --kubeconfig=/root/.kube/karmada.config --context=karmada-host
```

```console
kubectl get ns --show-labels --kubeconfig=/root/.kube/karmada.config --context=karmada-apiserver

kubectl apply -f karmada/ClusterOverridePolicy --kubeconfig=/root/.kube/karmada.config --context=karmada-apiserver

kubectl get ns --show-labels --kubeconfig=/root/.kube/members.config --context=member1
kubectl get ns --show-labels --kubeconfig=/root/.kube/members.config --context=member2
```