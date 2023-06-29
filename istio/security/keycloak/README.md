# [WIP] API Authentication using Istio Ingress Gateway, OAuth2-Proxy and Keycloak

Orignal post: https://medium.com/@senthilrch/api-authentication-using-istio-ingress-gateway-oauth2-proxy-and-keycloak-a980c996c259

## Cert

```console
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=zirain.info Inc./CN=zirain.info' -keyout zirain.info.root.key -out zirain.info.root.crt 
openssl req -out zirain.info.csr -newkey rsa:2048 -nodes -keyout zirain.info.key -subj "/CN=zirain.info/O=myexample organization"
openssl x509 -req -days 365 -CA zirain.info.root.crt -CAkey zirain.info.root.key -set_serial 0 -in zirain.info.csr -out zirain.info.crt
kubectl create -n istio-system secret tls zirain-info-credential --key=zirain.info.key --cert=zirain.info.crt
```

## Keycloak

```
kubectl apply -f https://raw.githubusercontent.com/keycloak/keycloak-quickstarts/latest/kubernetes-examples/keycloak.yaml -n istio-system
```

## Oath2-proxy

### Helm repo

```console
helm repo add oauth2-proxy https://oauth2-proxy.github.io/manifests
helm repo update
```

### Install

```console
helm upgrade --install --namespace istio-system --create-namespace \
  --values istio/security/ext-authz/oauth2-proxy.yaml oauth2-proxy oauth2-proxy/oauth2-proxy
```