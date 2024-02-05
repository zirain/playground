# EKS with X-Ray

This will not working, see more details [here](https://github.com/istio/istio/issues/36599#issuecomment-1744600462).

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


## Enable ADOT addon

### Install cert-manager

```shell
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.2/cert-manager.yaml
```

### Install ADOT

```shell
aws eks describe-addon-configuration --addon-name adot --addon-version v0.90.0-eksbuild.1
```

```shell
aws eks create-addon --addon-name adot --addon-version v0.90.0-eksbuild.1 --cluster-name appmeshtest-zirain
```

Verify ADOT status:

```shell
aws eks describe-addon --addon-name adot --cluster-name appmeshtest-zirain
```

Install Collector:

```shell
kubectl create ns observability
kubectl apply -f otel-collector.yaml
```

## Install Sample APP

```yaml
kubectl label namespace default istio-injection=enabled --overwrite
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.20/samples/sleep/sleep.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.20/samples/httpbin/httpbin.yaml
```


## Cleanup

```shell
kubectl delete otelcol --all -A
kubectl delete ns observability
aws eks delete-addon --addon-name adot --cluster-name appmeshtest-zirain
kubectl delete -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.2/cert-manager.yaml
aws eks delete-addon --addon-name adot --cluster-name appmeshtest-zirain
```