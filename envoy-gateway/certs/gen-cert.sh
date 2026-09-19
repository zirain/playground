#!/bin/bash
#
# (Re)generates every CA/leaf certificate used by this repo's TLS demos and
# syncs the matching Kubernetes Secrets/ConfigMaps. Safe to re-run: cert
# material and cluster objects are regenerated and replaced in place.
#
# Usage: ./gen-cert.sh

set -euo pipefail
cd "$(dirname "${BASH_SOURCE[0]}")"

CERT_DAYS=365

# Self-signed root CA: writes <name>.key / <name>.crt.
gen_ca() {
  local name=$1 cn=$2
  rm -f "${name}.srl"
  openssl req -x509 -sha256 -nodes -days "$CERT_DAYS" -newkey rsa:2048 \
    -subj "/O=example Inc./CN=${cn}" \
    -keyout "${name}.key" -out "${name}.crt"
}

# Leaf cert for $cn (used as both CN and SAN), signed by a CA made with
# gen_ca. Serials are tracked per-CA via <ca>.srl, so one CA can safely sign
# more than one leaf.
gen_leaf() {
  local name=$1 cn=$2 ca=$3
  openssl req -new -newkey rsa:2048 -nodes \
    -keyout "${name}.key" -out "${name}.csr" \
    -subj "/CN=${cn}/O=example organization"
  openssl x509 -req -days "$CERT_DAYS" \
    -CA "${ca}.crt" -CAkey "${ca}.key" -CAcreateserial \
    -in "${name}.csr" -out "${name}.crt" \
    -extfile <(printf 'keyUsage=keyEncipherment,digitalSignature\nextendedKeyUsage=serverAuth\nsubjectAltName=DNS:%s\n' "$cn")
}

# Creates/updates a kubernetes.io/tls Secret from a cert/key pair.
sync_tls_secret() {
  local secret=$1 cert=$2 key=$3 ns=${4:-}
  kubectl create secret tls "$secret" ${ns:+-n "$ns"} --key="$key" --cert="$cert" \
    --dry-run=client -o yaml | kubectl apply -f -
}

# Creates/updates a ConfigMap holding a CA cert under the `ca.crt` key, as
# required by BackendTLSPolicy's caCertificateRefs.
sync_ca_configmap() {
  local cm=$1 ca_crt=$2 ns=${3:-}
  kubectl create configmap "$cm" ${ns:+-n "$ns"} --from-file=ca.crt="$ca_crt" \
    --dry-run=client -o yaml | kubectl apply -f -
}

# example.com CA: www.example.com (httproute/backend-tls.yaml) and
# tls-backend-1 (httproute/mirror/tls-backend.yaml).
gen_ca example.com example.com
gen_leaf www.example.com www.example.com example.com
gen_leaf tls-backend-1.example.com tls-backend-1.example.com example.com

sync_tls_secret example-cert www.example.com.crt www.example.com.key
sync_ca_configmap example-ca example.com.crt
sync_tls_secret tls-backend-1-cert tls-backend-1.example.com.crt tls-backend-1.example.com.key

# example.org CA: www.example.org and the mirror-tls-backend demo's mirror
# target (a different issuer from tls-backend-1, for the mirror-to-different-
# cert demo).
gen_ca example.org example.org
gen_leaf www.example.org www.example.org example.org
gen_leaf mirror.example.com mirror.example.com example.org

sync_tls_secret example-org-cert www.example.org.crt www.example.org.key
sync_ca_configmap example-org-ca example.org.crt
sync_tls_secret mirror-tls-backend-cert mirror.example.com.crt mirror.example.com.key

# Client identity Envoy Gateway presents when originating backend mTLS
# (EnvoyProxy/backend-mtls/*).
gen_ca clientca example-client-ca
gen_leaf client example-client clientca

sync_tls_secret example-client-cert client.crt client.key envoy-gateway-system
sync_ca_configmap example-client-ca clientca.crt
