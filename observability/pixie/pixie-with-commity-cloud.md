# Community Cloud for Pixie


## Setup local cluster with k0s

```console
# Install k0s as a service
sudo k0s install controller --single

# Start k0s as a service
sudo k0s start

# Check service, logs and k0s status
sudo k0s status

# Export Admin's Kubeconfig file
sudo k0s kubeconfig admin > /root/.kube/config

# Check cluster status
kubectl get po -A
```

## Install Pixie

```console
helm install pixie pixie-operator/pixie-operator-chart --set deployKey="<px-dep-*>" --set clusterName=default-cluster --namespace pl --create-namespace
```


## Cleanup

```console
sudo k0s stop
sudo k0s reset
```