# istio network policy


## install kind with calico

```console
kind create cluster --image=kindest/node:v1.23.4 --config=kind-calico.yaml

kubectl apply -f https://docs.projectcalico.org/v3.23/manifests/calico.yaml

kubectl -n kube-system set env daemonset/calico-node FELIX_IGNORELOOSERPF=true


kubectl label namespace default istio-injection=enabled --overwrite
```

ic install --set profile=minimal --set values.global.logging.level="default:debug" --set meshConfig.accessLogFile=/dev/stdout --set values.global.proxy.holdApplicationUntilProxyStarts=true

curl http://169.254.169.254/latest/meta-data/instance-id