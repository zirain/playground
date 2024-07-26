#!/usr/bin/env bash

set -euo pipefail

kubectl apply -f ./k8s/kind/
kubectl wait --for=condition=ready pod --all --timeout=300s
kubectl get svc httpbin -oyaml
kubectl exec -it deploy/sleep -- curl -v httpbin:8000/get