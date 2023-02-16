```
istioctl install -f /root/go/src/github.com/zirain/playground/istio/iop/custom-tracing.yaml -y

kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/bookinfo/platform/kube/bookinfo.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/sleep/sleep.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/addons/jaeger.yaml

istioctl dashboard jaeger
```

kubectl exec -it sleep-78ff5975c6-bgmnx -- curl productpage:9080/productpage > /dev/null