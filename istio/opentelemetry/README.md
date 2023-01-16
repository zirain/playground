# Istio on OpenTelemetry

## Install Istio

```
istioctl install -f /root/go/src/github.com/zirain/playground/istio/opentelemetry/istio-operator.yaml -y
```

## Install OTel-collector

```
kubectl apply -f /root/go/src/github.com/zirain/playground/istio/opentelemetry/otel.yaml -nistio-system
```

```
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/addons/jaeger.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/addons/grafana.yaml


kubectl apply -f /root/go/src/github.com/zirain/playground/istio/opentelemetry/telemetry.yaml
```


kubectl label namespace default istio-injection=enabled --overwrite

kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/bookinfo/platform/kube/bookinfo.yaml


kubectl exec -it $(kubectl get po | grep sleep | awk '{print $1}') -- curl productpage:9080/productpage

kubectl port-forward -nistio-system pod/$(kubectl get pod -nistio-system | grep jaeger | awk '{print $1}') 8080:16686 --address 0.0.0.0