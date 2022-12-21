

```
helm repo add kube2iam https://jtblin.github.io/kube2iam/
helm install kube2iam kube2iam/kube2iam --set=extraArgs.base-role-arn=arn:aws:iam::017975021908:role/,extraArgs.default-role=kube2iam-default,host.iptables=true,host.interface=cali+ --set=rbac.create=true
```