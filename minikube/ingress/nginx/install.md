# install minikube
```
minikube start --image-mirror-country cn -registry-mirror=https://registry.docker-cn.com
```

# enable ingress addon

```
minikube addons enable ingress
```

# apply resource

```
kubectl apply -f helloworld.yml
```