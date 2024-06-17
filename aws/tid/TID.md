# TID

## Create EKS

```console
eksctl create cluster \
    --region=us-east-2 \
    --version 1.29 \
    --name tidtest-zirain \
    --nodes-min 2 \
    --nodes-max 3 \
    --nodes 2 \
    --auto-kubeconfig \
    --full-ecr-access 
```

```shell
export KUBECONFIG=~/.kube/eksctl/clusters/tidtest-zirain
```

## Install TID

```console
aws eks create-addon --addon-name tetrate-io_istio-distro \
    --addon-version v1.20.6-eksbuild.1 \
    --region=us-east-2 \
    --cluster tidtest-zirain
```

```console
aws eks describe-addon --cluster-name tidtest-zirain --addon-name tetrate-io_istio-distro
```


## Cleanup

```console
aws eks delete-addon --addon-name tetrate-io_istio-distro \
    --region=us-east-2 \
    --cluster tidtest-zirain

eksctl delete cluster --name tidtest-zirain --region=us-east-2
```