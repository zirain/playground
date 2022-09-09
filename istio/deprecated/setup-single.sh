#!/bin/bash

echo "begin to create cluster with kind"
kind create cluster --name istio
kubectl cluster-info --context kind-istio

echo "install metalb"
bash ../metalb/install-metallb.sh

echo "instal istio"
istioctl install --set profile=minimal -y

echo "finised!!!"