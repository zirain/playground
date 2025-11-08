
# Envoy Gateway GRPC demo


```shell
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=example Inc./CN=example.com' -keyout example.com.key -out example.com.crt

openssl req -out www.example.com.csr -newkey rsa:2048 -nodes -keyout www.example.com.key -subj "/CN=www.example.com/O=example organization"
openssl x509 -req -days 365 -CA example.com.crt -CAkey example.com.key -set_serial 0 -in www.example.com.csr -out www.example.com.crt

kubectl create secret tls example-cert --key=www.example.com.key --cert=www.example.com.crt

kubectl create configmap example-ca --from-file=example.com.crt
```


```shell
export GATEWAY_HOST=$(kubectl get gateway/example-gateway -o jsonpath='{.status.addresses[0].value}')
grpcurl -plaintext -authority=www.example.com ${GATEWAY_HOST}:80 yages.Echo/Ping

grpcurl --insecure -authority=www.example.com ${GATEWAY_HOST}:443 yages.Echo/Ping

grpcurl -insecure -authority=www.example.com ${GATEWAY_HOST}:4000 yages.Echo/Ping
```