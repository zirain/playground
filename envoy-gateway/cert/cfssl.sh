#!/bin/bash

# cfssl gencert -initca ca-csr.json | cfssljson -bare ca
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem \
     --config=ca-config.json -profile=envoygateway \
     gateway-csr.json | cfssljson -bare gateway

rm -rf gateway && mkdir gateway
mv gateway-key.pem gateway/tls.key
mv gateway.pem gateway/tls.crt
cp ca.pem gateway/ca.crt

kubectl delete secret envoy-gateway -n envoy-gateway-system --ignore-not-found
kubectl create secret generic envoy-gateway -n envoy-gateway-system \
      --from-file=gateway/tls.key \
      --from-file=gateway/tls.crt \
      --from-file=gateway/ca.crt
# rm -rf gateway

echo "Create envoy secret"
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem \
     --config=ca-config.json -profile=envoygateway \
     envoy-csr.json | cfssljson -bare envoy
     
rm -rf envoy && mkdir envoy
mv envoy-key.pem envoy/tls.key
mv envoy.pem envoy/tls.crt
cp ca.pem envoy/ca.crt

kubectl delete secret envoy -n envoy-gateway-system --ignore-not-found
kubectl create secret generic envoy -n envoy-gateway-system \
      --from-file=envoy/tls.key \
      --from-file=envoy/tls.crt \
      --from-file=envoy/ca.crt
# rm -rf envoy

echo "Create rate limit secret"

cfssl gencert -ca=ca.pem -ca-key=ca-key.pem \
     --config=ca-config.json -profile=envoygateway \
     envoy-csr.json | cfssljson -bare ratelimit

rm -rf ratelimit && mkdir ratelimit
mv ratelimit-key.pem ratelimit/tls.key
mv ratelimit.pem ratelimit/tls.crt
cp ca.pem ratelimit/ca.crt

kubectl delete secret envoy-rate-limit -n envoy-gateway-system --ignore-not-found
kubectl create secret generic envoy-rate-limit -n envoy-gateway-system \
      --from-file=ratelimit/tls.key \
      --from-file=ratelimit/tls.crt \
      --from-file=ratelimit/ca.crt
# rm -rf ratelimit
