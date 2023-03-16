# Istio gRPC demo

https://github.com/GoogleCloudPlatform/istio-samples/tree/master/sample-apps/grpc-greeter-go



```
istioctl install -y

kubectl label namespace default istio-injection=enabled

kubectl apply -f /root/go/src/github.com/zirain/playground/istio/networking/grpc/grpc-echo.yaml
```