

```
kubectl port-forward pod/envoy-gateway-conformance-infra-same-namespace-67617465-68vvt66 -n envoy-gateway-system 19000:19000
curl -X POST 127.0.0.1:19000/logging?level=debug

k logs envoy-gateway-conformance-infra-same-namespace-67617465-68vvt66 -n envoy-gateway-system

egctl config proxy c envoy-gateway-conformance-infra-same-namespace-67617465-68vvt66 -n envoy-gateway-system -oyaml


curl -v http://$GATEWAY_HOST/get -H "Host: www.example.com"
```