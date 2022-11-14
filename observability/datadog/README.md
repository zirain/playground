# DataDog


## Install datadog agent

```
kubectl create ns dg

# install operator
helm repo add datadog https://helm.datadoghq.com
helm install my-datadog-operator datadog/datadog-operator -n dg

# create secret
kubectl create secret generic datadog-secret -n dg --from-literal api-key=<YOUR_API_KEY> --from-literal app-key=<YOUR_APP_KEY> 

# apply agent resource
kubectl apply -f observability/datadog/agent.yaml -n dg

# install istio
istioctl install -f observability/datadog/iop.yaml --set hub=ghcr.io/zirain --set tag=1.17-dev -y

```
