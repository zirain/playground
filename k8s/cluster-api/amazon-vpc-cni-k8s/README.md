# Amazon VPC CNI

```console
clusterctl get kubeconfig capi-quickstart > /root/.kube/capi-quickstart.kubeconfig
kubectl apply -f https://raw.githubusercontent.com/aws/amazon-vpc-cni-k8s/v1.12.0/config/master/aws-k8s-cni.yaml --kubeconfig /root/.kube/capi-quickstart.kubeconfig
```

***注意放通master2node，node2node之间的安全组（包含TCP和UDP）***

```
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/httpbin/httpbin.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/sleep/sleep.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/helloworld/helloworld.yaml
```