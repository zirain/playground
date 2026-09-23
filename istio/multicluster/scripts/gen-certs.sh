#!/usr/bin/env bash

# Generate a shared root CA and one intermediate CA per cluster, so that
# workloads in different primaries trust each other.
# Usage: gen-certs.sh <cluster> [<cluster>...]

set -o errexit
set -o nounset
set -o pipefail

CERTS_DIR=${CERTS_DIR:-".certs"}
TRUST_DOMAIN=${TRUST_DOMAIN:-"cluster.local"}
mkdir -p "${CERTS_DIR}"

if [ ! -f "${CERTS_DIR}/root-key.pem" ]; then
  echo "Generating root CA"
  openssl genrsa -out "${CERTS_DIR}/root-key.pem" 4096
  openssl req -x509 -new -sha256 -days 3650 -key "${CERTS_DIR}/root-key.pem" \
    -subj "/O=Istio/CN=Root CA" \
    -addext "basicConstraints=critical,CA:true" \
    -addext "keyUsage=critical,digitalSignature,nonRepudiation,keyEncipherment,keyCertSign" \
    -out "${CERTS_DIR}/root-cert.pem"
fi

for cluster in "$@"; do
  dir="${CERTS_DIR}/${cluster}"
  [ -f "${dir}/ca-cert.pem" ] && continue
  echo "Generating intermediate CA for ${cluster}"
  mkdir -p "${dir}"
  openssl genrsa -out "${dir}/ca-key.pem" 4096
  openssl req -new -sha256 -key "${dir}/ca-key.pem" \
    -subj "/O=Istio/CN=Intermediate CA/L=${cluster}" \
    -out "${dir}/ca.csr"
  cat > "${dir}/ca.ext" <<EXT
basicConstraints=critical,CA:true,pathlen:0
keyUsage=critical,digitalSignature,nonRepudiation,keyEncipherment,keyCertSign
subjectAltName=URI:spiffe://${TRUST_DOMAIN}/ns/istio-system/sa/citadel
EXT
  openssl x509 -req -sha256 -days 730 -in "${dir}/ca.csr" \
    -CA "${CERTS_DIR}/root-cert.pem" -CAkey "${CERTS_DIR}/root-key.pem" -CAcreateserial \
    -extfile "${dir}/ca.ext" -out "${dir}/ca-cert.pem"
  cp "${CERTS_DIR}/root-cert.pem" "${dir}/root-cert.pem"
  cat "${dir}/ca-cert.pem" "${CERTS_DIR}/root-cert.pem" > "${dir}/cert-chain.pem"
done
