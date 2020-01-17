# Install

```
choco install kubernetes-helm
brew install kubernetes-helm
```

# Init

```
helm init --upgrade -i registry.cn-hangzhou.aliyuncs.com/google_containers/tiller:v2.14.2 --stable-repo-url https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts

helm init --upgrade --tiller-image gcr.azk8s.cn/kubernetes-helm/tiller:v2.14.2 --stable-repo-url https://mirror.azure.cn/kubernetes/charts/ --service-account tiller

```

# Reset

```
helm reset
helm reset -f
helm reset --remove-helm-home
```