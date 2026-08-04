# cert-manager + Vault demo

This demo shows a minimal cert-manager issuer that requests certificates from Vault using AppRole authentication.

## Prerequisites

- A running Kubernetes cluster
- kubectl configured to the cluster
- Helm installed
- A Vault server reachable from the cluster
- Permission to create namespaces, secrets, and issuers in the cluster

## 1. Install cert-manager
p
```bash
kubectl create namespace cert-manager
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm install cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --set installCRDs=true
```

Verify the installation:

```bash
kubectl get pods -n cert-manager
kubectl get crd | grep cert-manager
```

## 2. Install Vault

If you do not already have Vault, you can run a simple dev server locally for testing:

```bash
helm repo add hashicorp https://helm.releases.hashicorp.com
helm repo update
helm install vault hashicorp/vault \
  --namespace vault \
  --create-namespace \
  --set server.dev.enabled=true
```

For a production-like setup, expose Vault behind a reachable ingress or load balancer and use TLS.

## 3. Configure Vault PKI

Run the following commands inside Vault to enable a PKI secrets engine and issue a CA:

```bash
vault secrets enable pki
vault secrets tune -max-lease-ttl=8760h pki
vault write -field=certificate pki/root/generate/internal \
  common_name="example.com" \
  ttl=8760h > ca.pem
vault write pki/config/urls \
  issuing_certificates="http://vault.vault.svc.cluster.local:8200/v1/pki/ca" \
  crl_distribution_points="http://vault.vault.svc.cluster.local:8200/v1/pki/crl"
```

Create a role for cert-manager:

```bash
vault write pki/roles/example-dot-com \
  allowed_domains="example.com" \
  allow_subdomains=true \
  max_ttl="2160h"
vault write pki/roles/example1-dot-com \
  allowed_domains="example1.com" \
  allow_subdomains=true \
  max_ttl="2160h"
```

## 4. Create Vault AppRole

```bash
vault auth enable approle
vault write auth/approle/role/cert-manager-role \
  token_policies="cert-manager-demo" \
  secret_id_ttl=720h \
  token_ttl=20m \
  token_max_ttl=30m
vault write -f auth/approle/role/cert-manager-role/secret-id
vault read auth/approle/role/cert-manager-role/role-id
```

Create a policy that allows issuing certificates from the target PKI path. When Vault is running inside Kubernetes, copy the policy file into the Vault pod first:

```bash
kubectl cp ./vault-policy.hcl vault/vault-0:/tmp/vault-policy.hcl
kubectl exec -n vault vault-0 -- vault policy write cert-manager-demo /tmp/vault-policy.hcl
```

Make sure the AppRole has permission to issue certificates from the selected PKI path.

## 5. Prepare the Kubernetes secret

Create the Kubernetes secret with the AppRole values:

```bash
kubectl create secret generic vault-approle-credentials \
  -n cert-vault-demo \
  --from-literal=roleId='bbdfc669-5cbc-bb82-4adc-6ab590e82425' \
  --from-literal=secretId='ffac38ad-027b-58f3-c6c1-5d8d987ac1bb'
```

## 6. Update the demo manifests

Before applying the manifests, replace the following values:

- `https://vault.example.com` with your Vault server URL
- `pki/issuer/example-dot-com` with the Vault PKI signing path you want to use
- `tenant-a` with the Vault namespace, if required
- `replace-me` values in the Secret and Issuer with your real AppRole role ID and Vault secret ID
- `demo.example.com` with the DNS name you want to request

## 7. Apply the demo

```bash
kubectl apply -f namespace.yaml
kubectl apply -f secret.yaml
kubectl apply -f issuer.yaml
kubectl apply -f certificate.yaml
```

## 8. Verify the original demo

```bash
kubectl get issuer -n cert-vault-demo
kubectl get certificate -n cert-vault-demo
kubectl describe certificate demo-cert -n cert-vault-demo
kubectl get secret demo-cert-tls -n cert-vault-demo
```

If the issuer becomes Ready and the certificate is issued, the TLS secret will be created in the namespace.

## 9. Try a second issuer path without changing the existing one

This creates a fresh issuer and certificate so you can test a new issuance path while keeping the original resources intact.

```bash
kubectl apply -f issuer-step2.yaml
kubectl apply -f certificate-step2.yaml
```

Check the new issuance:

```bash
kubectl get issuer -n cert-vault-demo
kubectl get certificate -n cert-vault-demo
kubectl describe certificate demo-cert-step2 -n cert-vault-demo
kubectl get secret demo-cert-step2-tls -n cert-vault-demo
```

## What this demo includes

- A namespace for the demo resources
- A Kubernetes Secret holding the Vault AppRole credentials
- A cert-manager Issuer that talks to Vault
- A Certificate resource that requests a TLS secret


```shell
kubectl annotate certificate demo-cert -n cert-vault-demo cert-manager.io/renew-at="$(date -u +'%Y-%m-%dT%H:%M:%SZ')" --overwrite
```