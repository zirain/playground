#!/bin/bash

export TAG=1.12.0-debug
export TARGET_OS=linux

pushd ~/Codes/istio.io/istio
    make build-linux
popd
cp ~/Codes/istio.io/istio/out/linux_amd64/pilot-discovery .
sha256sum pilot-discovery
docker build -t istio/pilot:$TAG -f dockerfile.pilot .
kind load docker-image istio/pilot:$TAG --name istio
#kubectl rollout restart deployment/istiod -nistio-system
kubectl delete rs --all -nistio-system 
rm -rf pilot-discovery

docker image prune -y