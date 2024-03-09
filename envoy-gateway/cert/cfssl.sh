#!/bin/bash

UPDATE_SECRET=${1:-false}

cfssl gencert -initca config/ca-csr.json | cfssljson -bare ca

echo "Gen gateway cert"
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem \
     --config=config/ca-config.json -profile=cert \
     config/gateway-csr.json | cfssljson -bare gateway

rm -rf gateway && mkdir gateway
mv gateway-key.pem gateway/tls.key
mv gateway.pem gateway/tls.crt
cp ca.pem gateway/ca.crt

echo "Gen envoy cert"
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem \
     --config=config/ca-config.json -profile=cert \
     config/envoy-csr.json | cfssljson -bare envoy
     
rm -rf envoy && mkdir envoy
mv envoy-key.pem envoy/tls.key
mv envoy.pem envoy/tls.crt
cp ca.pem envoy/ca.crt

echo "Gen rate limit cert"

cfssl gencert -ca=ca.pem -ca-key=ca-key.pem \
     --config=config/ca-config.json -profile=cert \
     config/envoy-csr.json | cfssljson -bare ratelimit

rm -rf ratelimit && mkdir ratelimit
mv ratelimit-key.pem ratelimit/tls.key
mv ratelimit.pem ratelimit/tls.crt
cp ca.pem ratelimit/ca.crt

echo "Move certs to out folder"
rm -rf out && mkdir out
mv *pem out
mv *.csr out
mv envoy out
mv gateway out
mv ratelimit out

if $UPDATE_SECRET; then
  echo "Create gateway secret"
  kubectl delete secret envoy-gateway -n envoy-gateway-system --ignore-not-found
  kubectl create secret generic envoy-gateway -n envoy-gateway-system \
      --from-file=out/gateway/tls.key \
      --from-file=out/gateway/tls.crt \
      --from-file=out/gateway/ca.crt

  echo "Create envoy secret"
  kubectl delete secret envoy -n envoy-gateway-system --ignore-not-found
  kubectl create secret generic envoy -n envoy-gateway-system \
      --from-file=out/envoy/tls.key \
      --from-file=out/envoy/tls.crt \
      --from-file=out/envoy/ca.crt
  echo "Create rate limit secret"

  kubectl delete secret envoy-rate-limit -n envoy-gateway-system --ignore-not-found
  kubectl create secret generic envoy-rate-limit -n envoy-gateway-system \
      --from-file=out/ratelimit/tls.key \
      --from-file=out/ratelimit/tls.crt \
      --from-file=out/ratelimit/ca.crt
fi
