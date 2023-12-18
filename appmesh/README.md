# App Mesh with EKS

1. AWS Cli
1. eksctl
1. helm
    ```shell
    helm repo add eks https://aws.github.io/eks-charts
    helm repo update
    ```
1. Get AWS account id
   ```shell
   export AWS_ACCOUNT_ID=$(aws sts get-caller-identity | jq -r .Account)
   export AWS_DEFAULT_REGION=us-east-2
   ```
1. 


## Create EKS cluster

```shell
eksctl create cluster \
--name appmeshtest-zirain \
--nodes-min 2 \
--nodes-max 3 \
--nodes 2 \
--auto-kubeconfig \
--full-ecr-access \
--appmesh-access
```

```shell
export VPC_ID=$(aws eks describe-cluster --name appmeshtest-zirain | jq -r .cluster.resourcesVpcConfig.vpcId)
```

## Setup appmesh controller

```shell
kubectl apply -k "https://github.com/aws/eks-charts/stable/appmesh-controller/crds?ref=master"
```

```shell
eksctl utils associate-iam-oidc-provider \
    --region=us-east-2 \
    --cluster appmeshtest-zirain \
    --approve
```

> Note: if you deleted serviceaccount in the appmesh-system namespace, you will need to delete and re-create iamserviceaccount. eksctl does not override the iamserviceaccount correctly [see this issue](https://github.com/eksctl-io/eksctl/issues/2665)

```shell
eksctl create iamserviceaccount \
    --cluster appmeshtest-zirain \
    --namespace appmesh-system \
    --region=us-east-2 \
    --name appmesh-controller \
    --attach-policy-arn  arn:aws:iam::aws:policy/AWSCloudMapFullAccess,arn:aws:iam::aws:policy/AWSAppMeshFullAccess \
    --override-existing-serviceaccounts \
    --approve
```

```shell
helm upgrade -i appmesh-controller eks/appmesh-controller \
    --namespace appmesh-system \
    --set region=us-east-2 \
    --set serviceAccount.create=false \
    --set serviceAccount.name=appmesh-controller
```

## Cleanup

```shell
eksctl delete iamserviceaccount --cluster appmeshtest-zirain --namespace appmesh-system --region=us-east-2 --name appmesh-controller
helm delete appmesh-controller -n appmesh-system
eksctl delete cluster --name appmeshtest-zirain --region=us-east-2
```