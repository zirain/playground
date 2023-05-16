
```
kind export kubeconfig --name kurator-member1 --kubeconfig=/root/.kube/kurator-member1.config --internal
kind export kubeconfig --name kurator-member2 --kubeconfig=/root/.kube/kurator-member2.config --internal
```

```
kubectl delete secret kurator-member1 kurator-member2

kubectl create secret generic kurator-member1 --from-file=/root/.kube/kurator-member1.config
kubectl create secret generic kurator-member2 --from-file=/root/.kube/kurator-member2.config
```

```
kubectl rollout restart deploy/kurator-fleet-manager -n kurator-system

kubectl logs -n kurator-system -l app.kubernetes.io/name=kurator-fleet-manager --tail=-1 -f

```