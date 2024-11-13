
## Passthrough

```shell
kubectl apply -f istio/networking/tls-demo/tls-passthrough.yaml
curl -ik -v -HHost:passthrough.example.com --resolve "passthrough.example.com:443:172.18.0.205" https://passthrough.example.com:443/get
```

## Simple

```shell
kubectl apply -f istio/networking/tls-demo/tls-simple.yaml

curl -ik -v -HHost:v1.example.com --resolve "v1.example.com:443:172.18.0.205" https://v1.example.com:443/get

curl -ik -v -HHost:v2.example.com --resolve "v2.example.com:443:172.18.0.205" https://v2.example.com:443/get
```