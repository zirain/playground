# Skywalking


```console
kubectl create ns skywalking
```

## Setup elk

```console
kubectl apply -f elk/elk.yaml -nskywalking
```

## Install skywalking

```
helm repo add skywalking https://apache.jfrog.io/artifactory/skywalking-helm

```

```console
helm repo add skywalking https://apache.jfrog.io/artifactory/skywalking-helm

helm install skywalking skywalking/skywalking -n skywalking -f istio/telemetry/skywalking/skywalking.yaml

# port forward skywalking ui
kubectl port-forward svc/skywalking-ui 3000:80 --namespace skywalking --address 0.0.0.0

```


## Install istioctl

```console
istioctl -f istio/telemetry/skywalking/iop.yaml
```

enable mesh tracing:

```console
kubectl apply -f istio/telemetry/skywalking/iop.yaml
```


## Bookinfo

```console

kubectl label namespace default istio-injection=enabled --overwrite
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/bookinfo/platform/kube/bookinfo.yaml

```



## cleanup

```console
helm uninstall skywalking -n skywalking
```