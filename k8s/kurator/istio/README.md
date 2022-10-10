```console
kubectl label namespace default istio-injection=enabled --overwrite


kubectl apply --kubeconfig=/etc/karmada/karmada-apiserver.config -f networking/
```

label cluster for install istio with different networks:

```
kubectl label cluster member1 topology.istio.io/network=member1 --overwrite --kubeconfig=/etc/karmada/karmada-apiserver.config
kubectl label cluster member2 topology.istio.io/network=member2 --overwrite --kubeconfig=/etc/karmada/karmada-apiserver.config
```


```
LOGGING_LEVEL=debug out/linux-amd64/kurator install istio --primary member1 --remote member2 --set meshConfig.accessLogFile=/dev/stdout
```