# thanos 


## TODO

- 添加 external_labels https://github.com/prometheus-operator/prometheus-operator/issues/2650
- 确认可以通过Headless服务发现sidecar endpoint


```console
out/linux-amd64/kurator install thanos --host-kubeconfig /root/.kube/kurator-host.config --host-context kurator-host --object-store-config /root/thanos/thanos-config.yaml
```