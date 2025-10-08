https://gateway.envoyproxy.io/docs/tasks/security/backend-mtls/


```shell
kubectl create secret tls example-cert --key=www.example.com.key --cert=www.example.com.crt
kubectl create configmap example-ca --from-file=ca.crt
kubectl create secret -n envoy-gateway-system tls example-client-cert --key=client.key --cert=client.crt
kubectl create configmap example-client-ca --from-file=clientca.crt
```