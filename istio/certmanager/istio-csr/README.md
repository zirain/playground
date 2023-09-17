# Securing the istio Service Mesh using cert-manager

## Helm setup

```bash
# Helm setup
helm repo add jetstack https://charts.jetstack.io
helm repo update
```

## Installing cert-manager

``` bash
# install cert-manager; this might take a little time
helm install cert-manager jetstack/cert-manager \
	--namespace cert-manager \
	--set installCRDs=true \
	--create-namespace \
	--version v1.8.0

# We need this namespace to exist since our cert will be placed there
kubectl create namespace istio-system

# Create a cert-manager Issuer and Issuing Certificate
kubectl apply -f istio/certmanager/istio-csr/example-issuer.yaml
```

## Export the Root CA to a local File

```bash
# Export our cert from the secret it's stored in, and base64 decode to get the PEM data.
kubectl get -n istio-system secret istio-ca -ogo-template='{{index .data "tls.crt"}}' | base64 -d > root-ca.pem

# Out of interest, we can check out what our CA looks like
openssl x509 -in root-ca.pem -noout -text

# Add our CA to a secret
kubectl create secret generic -n cert-manager istio-root-ca --from-file=ca.pem=root-ca.pem
```

## Installing istio-csr

```bash
# We set a few helm template values so we can point at our static root CA
helm install -n cert-manager cert-manager-istio-csr jetstack/cert-manager-istio-csr -f istio/certmanager/istio-csr/example-issuer.yaml

kubectl rollout status deploy cert-manager-istio-csr -n cert-manager --timeout 5m

# Check to see that the istio-csr pod is running and ready
kubectl get pods -n cert-manager
NAME                                       READY   STATUS    RESTARTS   AGE
cert-manager-aaaaaaaaaa-11111              1/1     Running   0          9m46s
cert-manager-cainjector-aaaaaaaaaa-22222   1/1     Running   0          9m46s
cert-manager-istio-csr-bbbbbbbbbb-00000    1/1     Running   0          63s
cert-manager-webhook-aaaaaaaaa-33333       1/1     Running   0          9m46s
```

## Installing Istio

```bash
istioctl install -f istio/certmanager/istio-csr/istio-iop.yaml -y
```

```bash
kubectl logs -n cert-manager $(kubectl get pods -n cert-manager -o jsonpath='{.items..metadata.name}' --selector app=cert-manager) --since 2m -f

kubectl get certificaterequests.cert-manager.io -n istio-system -w

```