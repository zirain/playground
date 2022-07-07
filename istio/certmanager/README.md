# cert-manager CA 集成

# install

```console
helm repo add jetstack https://charts.jetstack.io
helm repo update
kubectl create namespace cert-manager
helm install -n cert-manager cert-manager jetstack/cert-manager --set installCRDs=true
```

```console
kubectl apply -f istio/certmanager/certmanger.yaml -nistio-system

# verify
kubectl get issuers -n istio-system
```

export our Root CA and configure Istio later using that static cert, see more details [here](https://cert-manager.io/docs/tutorials/istio-csr/istio-csr/#export-the-root-ca-to-a-local-file):
```
# Export our cert from the secret it's stored in, and base64 decode to get the PEM data.
kubectl get -n istio-system secret istio-ca -ogo-template='{{index .data "tls.crt"}}' | base64 -d > ca.pem

# Out of interest, we can check out what our CA looks like
openssl x509 -in ca.pem -noout -text

# Add our CA to a secret
kubectl create secret generic -n cert-manager istio-root-ca --from-file=ca.pem=ca.pem
```

install istio-csr
```console
helm install -n cert-manager cert-manager-istio-csr jetstack/cert-manager-istio-csr \
	--set "app.tls.rootCAFile=/var/run/secrets/istio-csr/ca.pem" \
	--set "volumeMounts[0].name=root-ca" \
	--set "volumeMounts[0].mountPath=/var/run/secrets/istio-csr" \
	--set "volumes[0].name=root-ca" \
	--set "volumes[0].secret.secretName=istio-root-ca"

kubectl get certificates -n istio-system
```

```console
istioctl install -f istio/certmanager/iop.yaml -y
```

```console
istioctl proxy-config secret $(kubectl get pods -o jsonpath='{.items..metadata.name}' --selector app=httpbin) -o json | jq -r '.dynamicActiveSecrets[0].secret.tlsCertificate.certificateChain.inlineBytes' | base64 --decode | openssl x509 -text -noout
```

## cleanup

```
istioctl x uninstall --purge -y

helm uninstall -n cert-manager cert-manager-istio-csr

kubectl delete -f istio/certmanager/certmanger.yaml -nistio-system

helm uninstall -n cert-manager cert-manager
```

## Reference

- https://docs.tetrate.io/zh/istio-ca-certs-integrations/cert-manager-integration/
- https://cert-manager.io/docs/projects/istio-csr/
