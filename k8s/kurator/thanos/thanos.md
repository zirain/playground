# thanos 


## TODO

- 添加 external_labels https://github.com/prometheus-operator/prometheus-operator/issues/2650
- 确认可以通过Headless服务发现sidecar endpoint

## Setup

```console
out/linux-amd64/kurator install thanos --host-kubeconfig /root/.kube/kurator-host.config --host-context kurator-host --object-store-config /root/thanos/thanos-config.yaml
```

```
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/addons/extras/prometheus-operator.yaml --kubeconfig $HOME/.kube/kurator-members.config --context kurator-member1
```


## Cleanup

```console
kubectl delete -f manifests/profiles/prom-thanos/ --ignore-not-found --kubeconfig /etc/karmada/karmada-apiserver.config
kubectl delete -f manifests/profiles/prom-thanos/setup --ignore-not-found --kubeconfig /etc/karmada/karmada-apiserver.config
kubectl delete ns monitoring --ignore-not-found --kubeconfig /etc/karmada/karmada-apiserver.config
kubectl delete cpp thanos --kubeconfig /etc/karmada/karmada-apiserver.config

kubectl delete -f manifests/profiles/thanos/ --ignore-not-found --kubeconfig $HOME/.kube/kurator-host.config
kubectl delete ns thanos --ignore-not-found --kubeconfig $HOME/.kube/kurator-host.config
```