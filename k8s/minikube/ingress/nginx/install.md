# install minikube
```shell
minikube start --image-mirror-country cn -registry-mirror=https://registry.docker-cn.com
```

# enable ingress addon

```shell
minikube addons enable ingress
```

# apply resource

```shell
kubectl apply -f helloworld.yml
```