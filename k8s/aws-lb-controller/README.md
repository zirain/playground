# AWS Load Balancer Controller

Project URL: https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.4/

Operator from opensift: https://github.com/openshift/aws-load-balancer-operator

## Quick start


根据 `iam-policy.json` 内容创建策略 `AWSLoadBalancerControllerIAMPolicy` 并绑定角色 `nodes.cluster-api-provider-aws.sigs.k8s.io`


aws load balancer controller严重依赖于资源的标签, cluster-api-provider-aws 创建时相关资源时会正确标记
因此启动参数 `cluster-name` 的值务必与 `cluster-api` 创建的集群名称相同

```
kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.5.4/cert-manager.yaml

kubectl apply -f v2_4_5_full.yaml
kubectl apply -f v2_4_5_ingclass.yaml
```


## Echo Demo

kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v2.4.5/docs/examples/echoservice/echoserver-namespace.yaml &&\
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v2.4.5/docs/examples/echoservice/echoserver-service.yaml &&\
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v2.4.5/docs/examples/echoservice/echoserver-deployment.yaml


kubectl apply -f echoserver-ingress.yaml

kubectl describe ing -n echoserver echoserver


curl  k8s-echoserv-echoserv-b9c0cf078f-1033584182.us-east-1.elb.amazonaws.com --header "Host: echoserver.example.com"

curl -v echoserver.example.com --resolve "echoserver.example.com:80:34.236.229.225"
