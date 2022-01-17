#!/bin/bash


kubectl label namespace default istio-injection=enabled
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.12/samples/httpbin/httpbin.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.12/samples/httpbin/sample-client/fortio-deploy.yaml 
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.12/sleep/sleep.yaml

