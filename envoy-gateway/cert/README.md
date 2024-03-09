# Custom ControlPlane Certificates


## cfssl

```shell
bash cfssl.sh
```

## Cert-manager

```shell
helm install cert-manager jetstack/cert-manager \
	--namespace cert-manager \
	--set installCRDs=true \
	--create-namespace \
	--version v1.14.4

kubectl apply -f cert-manager.yaml
```

## Verify

```shell
kubectl get secret -n envoy-gateway-system envoy -o jsonpath="{.data['ca\.crt']}"  | base64 -d - | step certificate inspect -
kubectl get secret -n envoy-gateway-system envoy -o jsonpath="{.data['tls\.crt']}" | base64 -d - | step certificate inspect -
kubectl get secret -n envoy-gateway-system envoy -o jsonpath="{.data['tls\.key']}" | base64 -d - | step certificate inspect -
```

```shell
kubectl get secret -n envoy-gateway-system envoy-gateway -o jsonpath="{.data['ca\.crt']}"  | base64 -d - | step certificate inspect -
kubectl get secret -n envoy-gateway-system envoy-gateway -o jsonpath="{.data['tls\.crt']}" | base64 -d - | step certificate inspect -
kubectl get secret -n envoy-gateway-system envoy-gateway -o jsonpath="{.data['tls\.key']}" | base64 -d - | step certificate inspect -
```

## References

* https://www.baeldung.com/openssl-self-signed-cert
* https://github.com/coreos/docs/blob/master/os/generate-self-signed-certificates.md
* https://smallstep.com/practical-zero-trust/go-grpc-tls
* https://smallstep.com/docs/step-cli/basic-crypto-operations/#requirements
* https://github.com/cloudflare/cfssl
* 