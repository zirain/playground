# AWS EFS CSI

## Create EFS

```
# 获取VPC ID
vpc_id=$(aws ec2 describe-vpcs --filter Name=tag:Name,Values=capi-quickstart-vpc --query Vpcs[].VpcId --output text)
# 检索VPC CIDR
cidr_range=$(aws ec2 describe-vpcs --vpc-ids $vpc_id --query "Vpcs[].CidrBlock" --output text)
# 创建 EFS 安全组
security_group_id=$(aws ec2 create-security-group --group-name MyEfsSecurityGroup --description "My EFS security group" --vpc-id $vpc_id --output text)
# 创建入站规则
aws ec2 authorize-security-group-ingress --group-id $security_group_id --protocol tcp --port 2049 --cidr $cidr_range
```

## Install wit helm

```
helm repo add aws-efs-csi-driver https://kubernetes-sigs.github.io/aws-efs-csi-driver/
helm repo update

# IRSA
helm upgrade --install aws-efs-csi-driver --namespace kube-system aws-efs-csi-driver/aws-efs-csi-driver

# kube2iam
helm upgrade --install aws-efs-csi-driver --namespace kube-system aws-efs-csi-driver/aws-efs-csi-driver -f /root/aws/efs-csi-values.yaml
```


## Examples

https://github.com/kubernetes-sigs/aws-efs-csi-driver/blob/master/examples/kubernetes/static_provisioning/README.md


## DEBUG

```console
kubectl logs -l app=efs-csi-controller -nkube-system --tail=-1

kubectl logs -l app=efs-csi-node -nkube-system --tail=-1
```
