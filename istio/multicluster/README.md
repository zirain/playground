# Istio Multicluster

## Primary remote 

### Same network

```shell
make create-clusters connect-clusters install-istio install-app
```


### Different networks

```shell
ISTIO_NETWORK_MODE=non-flat make create-clusters connect-clusters install-istio install-app
```


## Ambient multi-primary

### Different networks

Two primary clusters (`cluster1` on `network1`, `cluster2` on `network2`) running ambient mode,
connected through HBONE east-west gateways. A shared root CA is generated into `.certs`.

```shell
ISTIO_MC_MODE=ambient-multi-primary ISTIO_NETWORK_MODE=non-flat make up
```

Verify cross-cluster traffic (should return both `v1` and `v2`):

```shell
for i in $(seq 10); do
  kubectl --kubeconfig .kube/cluster1 exec -n ambient deploy/sleep -- curl -s helloworld:5000/hello
done
```

Services are shared across clusters (`serviceScopeConfigs` in the IOPs) when they are:
- in a namespace labeled `istio.io/dataplane-mode=ambient`, unless the service is labeled `istio.io/global=false`, or
- labeled `istio.io/global=true`, in any namespace.

## Cleanup

```shell
make down
```

## Tips

### How to install Istio multicluster manually

- Install istio in primary cluster
- Expose istiod service
- Create remote secret with `istioctl x create-remote-secret --server https://172.18.0.3:6443 --name remote > remote.kubeconfig`
  `172.18.0.3` is ip of the docker container
- Install remote istio with discovery address `kubectl get --kubeconfig=${MAIN_KUBECONFIG} svc -nistio-system istio-eastwestgateway  -o jsonpath="{.status.loadBalancer.ingress[0].ip}"`