
```
kind export kubeconfig --name kurator-member1 --kubeconfig=/root/.kube/kubrator-member1.config
kind export kubeconfig --name kurator-member2 --kubeconfig=/root/.kube/kubrator-member2.config
```

```
kubectl create secret generic kurator-member1 --from-file=/root/.kube/kubrator-member1.config
kubectl create secret generic kurator-member2 --from-file=/root/.kube/kubrator-member2.config
```