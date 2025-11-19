# DataDog


## Install datadog agent

```
kubectl create ns dg

# install operator
helm repo add datadog https://helm.datadoghq.com
helm install dg-operator datadog/datadog-operator -n datadog-agent

# create secret
kubectl create secret generic datadog-secret -n datadog-agent --from-literal api_key=<YOUR_API_KEY>

# apply agent resource
kubectl apply -f observability/datadog/agent.yaml -n datadog-agent

# install istio
istioctl install -f observability/datadog/iop.yaml --set hub=ghcr.io/zirain --set tag=1.17-dev -y

```
