# Cluster API AWS provider FAQ

## 支持 EKS VPC CNI (WIP)

   Project URL: https://github.com/aws/amazon-vpc-cni-k8s
   Install with helm: https://github.com/aws/eks-charts/tree/master/stable/aws-vpc-cni
   ```
   kubectl apply -f https://raw.githubusercontent.com/aws/amazon-vpc-cni-k8s/v1.12.0/config/master/aws-k8s-cni.yaml
   ```

## AK/SK 多租户支持

   https://cluster-api-aws.sigs.k8s.io/topics/multitenancy.html#awsclusterroleidentity

## 不同用户的secret存储安全

## 节点探活过程，集群状态维护逻辑，探活的方式是否能够拆分？

## EKS默认`ingress`/`ELB`方案是否为 `AWS Load Balancer Controller`？

   https://docs.aws.amazon.com/eks/latest/userguide/alb-ingress.html

## CSI 跟 S3/EVS/SFS 的集成，对齐EKS?

   [EBS](https://github.com/kubernetes-sigs/aws-ebs-csi-driver)
   [EFS](https://github.com/kubernetes-sigs/aws-efs-csi-driver)
   [FSx](https://github.com/kubernetes-sigs/aws-fsx-csi-driver)

## `Secret`加解密对接KMS

   https://github.com/kubernetes-sigs/aws-encryption-provider

## `ServiceAccount` 对接 IAM

   [kube2iam](https://github.com/jtblin/kube2iam)

## 节点OS image支持自定义，不同region镜像名是否一致？

    https://cluster-api-aws.sigs.k8s.io/topics/images/amis.html


    clusterctl get kubeconfig capi-multitenancy > /root/.kube/capi-quickstart.kubeconfig