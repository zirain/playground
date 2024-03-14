# Mutual TLS origination for egress traffic

1. Generate client and server certificates and keys

```shell
bash gen-ssl.sh
```

```shell
kubectl create namespace mesh-external

kubectl create -n mesh-external secret tls nginx-server-certs --key my-nginx.mesh-external.svc.cluster.local.key --cert my-nginx.mesh-external.svc.cluster.local.crt
kubectl create -n mesh-external secret generic nginx-ca-certs --from-file=example.com.crt

kubectl apply -f nginx-ssl.yaml -n mesh-external
```


```shell
kubectl create secret generic client-credential --from-file=tls.key=client.example.com.key \
  --from-file=tls.crt=client.example.com.crt --from-file=ca.crt=example.com.crt

kubectl create role client-credential-role --resource=secret --verb=list
kubectl create rolebinding client-credential-role-binding --role=client-credential-role --serviceaccount=default:sleep

kubectl exec -it deploy/sleep -c sleep -- curl -sS http://my-nginx.mesh-external.svc.cluster.local
```