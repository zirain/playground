# AWS Encryption Provider

https://github.com/kubernetes-sigs/aws-encryption-provider

```
REPO=ghcr.io/zirain make build-server build-docker
```

## Create KMS key

```
KEY_ID=$(aws --region us-east-1 kms create-key --description "kms for etcd" --query KeyMetadata.KeyId --output text)
aws kms describe-key --key-id $KEY_ID
```


```
aws iam create-policy --policy-name kms-policy --policy-document file:///root/go/src/github.com/zirain/playground/k8s/cluster-api/aws-encryption-provider/iam-policy.json

aws iam attach-role-policy --policy-arn arn:aws:iam::017975021908:policy/kms-policy --role-name control-plane.cluster-api-provider-aws.sigs.k8s.io
```


```
aws s3 cp /root/go/src/github.com/zirain/playground/k8s/cluster-api/aws-encryption-provider/aws-encryption-provider.yaml s3://kurator-operator-test/aws-encryption-provider.yaml
aws s3 cp /root/go/src/github.com/zirain/playground/k8s/cluster-api/aws-encryption-provider/encryption-provider-config.yaml s3://kurator-operator-test/encryption-provider-config.yaml
```


## Verify

```
kubectl apply -f /root/go/src/github.com/zirain/playground/k8s/cluster-api/aws-encryption-provider/capi-quickstart.yaml

clusterctl get kubeconfig capi-quickstart > /root/.kube/capi-quickstart.kubeconfig
kubectl --kubeconfig=/root/.kube/capi-quickstart.kubeconfig apply -f https://raw.githubusercontent.com/projectcalico/calico/v3.24.1/manifests/calico.yaml

kubectl create secret generic secret1 -n default --from-literal=mykey=mydata

wget https://github.com/etcd-io/etcd/releases/download/v3.4.23/etcd-v3.4.23-linux-amd64.tar.gz
tar -zxvf etcd-v3.4.23-linux-amd64.tar.gz
ETCDCTL_API=3 etcdctl \
    --key /etc/kubernetes/pki/etcd/server.key \
    --cert /etc/kubernetes/pki/etcd/server.crt \
    --cacert /etc/kubernetes/pki/etcd/ca.crt  \
    --endpoints "https://10.0.80.65:2379" get /registry/secrets/default/secret1

```