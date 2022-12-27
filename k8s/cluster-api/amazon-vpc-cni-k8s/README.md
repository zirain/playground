# Amazon VPC CNI

https://github.com/aws/amazon-vpc-cni-k8s

```
{
    "Effect": "Allow",
    "Principal": {
        "Federated": "arn:aws:iam::017975021908:oidc-provider/kurator-operator-test.s3.amazonaws.com"
    },
    "Action": "sts:AssumeRoleWithWebIdentity",
    "Condition": {
        "StringEquals": {
            "kurator-operator-test.s3.amazonaws.com:sub": "system:serviceaccount:kube-system:aws-load-balancer-controller"
        }
    }
}
```

```console
clusterctl get kubeconfig capi-quickstart > /root/.kube/capi-quickstart.kubeconfig
kubectl apply -f /root/go/src/github.com/zirain/playground/k8s/cluster-api/amazon-vpc-cni-k8s/aws-k8s-cni.yaml --kubeconfig /root/.kube/capi-quickstart.kubeconfig
```

***注意放通master2node，node2node之间的安全组（包含TCP和UDP）***

```
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/sleep/sleep.yaml --kubeconfig /root/.kube/capi-quickstart.kubeconfig
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/helloworld/helloworld.yaml --kubeconfig /root/.kube/capi-quickstart.kubeconfig
```