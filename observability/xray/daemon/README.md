# EKS with X-Ray Daemonset

## Create Cluster

```shell
eksctl create cluster \
    --region=us-east-2 \
    --name appmeshtest-zirain \
    --nodes-min 2 \
    --nodes-max 3 \
    --nodes 2 \
    --auto-kubeconfig \
    --full-ecr-access \
    --appmesh-access
```

```shell
export KUBECONFIG=~/.kube/eksctl/clusters/appmeshtest-zirain
```

## Install X-Ray Daemonset

```shell
eksctl create iamserviceaccount --name xray-daemon --namespace xray-system \
    --cluster appmeshtest-zirain \
    --attach-policy-arn arn:aws:iam::aws:policy/AWSXRayDaemonWriteAccess \
    --override-existing-serviceaccounts \
    --approve 

kubectl label serviceaccount xray-daemon app=xray-daemon --namespace xray-system
```

```shell
kubectl apply -f observability/xray/daemon/xray-k8s-daemonset.yaml --namespace xray-system
kubectl describe daemonset xray-daemon --namespace xray-system
```

## Cleanup

```shell
kubectl delete -f observability/xray/daemon/xray-k8s-daemonset.yaml --namespace xray-system
eksctl delete iamserviceaccount --name xray-daemon --namespace xray-system --cluster appmeshtest-zirain
eksctl delete cluster --region=us-east-2 --name appmeshtest-zirain
```