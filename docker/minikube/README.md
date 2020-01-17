# Install 

```
minikube start --registry-mirror=https://registry.docker-cn.com --vm-driver="hyperv" --hyperv-virtual-switch="minikubeSwitch"
```

# Use ACR in Minikube

## Create AAD Service Principals

```
az ad sp create-for-rbac
  --scopes /subscriptions/<SUBSCRIPTION_ID>/resourcegroups/<RG_NAME>/providers/Microsoft.ContainerRegistry/registries/<REGISTRY_NAME>
  --role Contributor
  --name <SERVICE_PRINCIPAL_NAME>
```


## Configure Kubernetes to use your ACR

```
kubectl create secret docker-registry <SECRET_NAME>
  --docker-server <REGISTRY_NAME>.azurecr.io
  --docker-email <YOUR_MAIL>
  --docker-username=<SERVICE_PRINCIPAL_ID>
  --docker-password <YOUR_PASSWORD>
```

## Create Pods, Replica Sets, and Deployments using the Secret

```
apiVersion: v1
kind: Pod
metadata:
  name: private-hello-world
spec:
  containers:
  - name: private-hello-container
    image: <REGISTRY_NAME>.azurecr.io/hello-world
  imagePullSecrets:
  - name: <SECRET_NAME>
```