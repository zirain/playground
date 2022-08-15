# Custom CNI configuration with k0s

## 

- Generate a yaml config file that uses the default settings.

```console
mkdir -p /etc/k0s
k0s config create > /etc/k0s/k0s.yaml
```

- Modify `spec.network.provider` to `custom`(default value is `kuberouter`)
- Add extesion with `cilium` chart
- Install k0s with your new config file

```console
sudo k0s install controller -c /etc/k0s/k0s.yaml --single
```

- Start k0s

```console
sudo k0s start

# Check service, logs and k0s status
sudo k0s status

# Export Admin's Kubeconfig file
sudo k0s kubeconfig admin > /root/.kube/config
```
- Install Metallb

```console
bash k8s/metallb/install-metalb.sh
```

- Expose hubble ui

```
kubectl apply -f k8s/cilium/hubble-ui.yaml
```

- Setup demo app

```
kubectl create -f https://raw.githubusercontent.com/cilium/cilium/1.12.0/examples/minikube/http-sw-app.yaml
```

- Check current access

```
kubectl exec xwing -- curl -s -XPOST deathstar.default.svc.cluster.local/v1/request-landing
kubectl exec tiefighter -- curl -s -XPOST deathstar.default.svc.cluster.local/v1/request-landing
```


## Cleanup

```
sudo k0s stop
sudo k0s reset
```