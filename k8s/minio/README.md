# Minio

```console
helm install minio oci://registry-1.docker.io/bitnamicharts/minio -n monitoring \
    -f /root/go/src/github.com/zirain/playground/k8s/minio/minio.values.yaml \
    --create-namespace

kubectl delete secret thanos-objstore
kubectl create secret generic thanos-objstore --from-file=objstore.yml=/root/go/src/github.com/zirain/playground/k8s/minio/obj.yaml

kubectl delete secret thanos-objstore -n monitoring --kubeconfig=/root/.kube/kurator-member1.config --ignore-not-found
kubectl create secret generic thanos-objstore --from-file=objstore.yml=/root/go/src/github.com/zirain/playground/k8s/minio/obj.yaml -n monitoring --kubeconfig=/root/.kube/kurator-member1.config

kubectl delete secret thanos-objstore -n monitoring --kubeconfig=/root/.kube/kurator-member2.config --ignore-not-found
kubectl create secret generic thanos-objstore --from-file=objstore.yml=/root/go/src/github.com/zirain/playground/k8s/minio/obj.yaml -n monitoring --kubeconfig=/root/.kube/kurator-member2.config
```