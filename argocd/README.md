# ArgoCD

```shell
kubectl -n argocd apply -f install.yaml

kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d

argocd login <lb_ip> --insecure


argocd account update-password --account adminuser

kubectl config get-clusters
argocd cluster add <your_cluster_context> --in-cluster --name in-cluster

argocd account generate-token
```