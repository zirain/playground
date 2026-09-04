

## Create a ServiceAccount & Bearer Token


```bash
kubectl create serviceaccount envoy-api-user -n default
kubectl create clusterrolebinding envoy-api-user-admin \
  --clusterrole=cluster-admin \
  --serviceaccount=default:envoy-api-user

# Generate a 1-year Bearer Token
API_TOKEN=$(kubectl create token envoy-api-user --duration=8760h -n default)
echo "Token generated successfully"
```


## Get the Cluster CA Certificate Secret

```bash
# Extract the cluster CA from the default serviceaccount token
kubectl get secret $(kubectl get serviceaccount default -o jsonpath='{.secrets[0].name}') \
  -o jsonpath='{.data.ca\.crt}' | base64 -d > /tmp/ca.crt 2>/dev/null || \
kubectl get cm kube-root-ca.crt -o jsonpath='{.data.ca\.crt}' > /tmp/ca.crt

# Create secret for Envoy Gateway BackendTLSPolicy
kubectl create secret generic apiserver-ca-cert \
  --from-file=ca.crt=/tmp/ca.crt \
  -n default
```

## Create Gateway Certificate (Frontend TLS)

```bash
# Generate CA and server key/cert for Envoy Gateway
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout /tmp/tls.key -out /tmp/tls.crt \
  -subj "/CN=k8s-api.example.com" \
  -addext "subjectAltName = DNS:k8s-api.example.com"

# Create TLS secret for Gateway Frontend
kubectl create secret tls gateway-frontend-tls \
  --key=/tmp/tls.key \
  --cert=/tmp/tls.crt \
  -n default
```