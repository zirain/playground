#!/usr/bin/env bash

set -euo pipefail


function dump_logs() {
  kubectl logs -l app=whoami -c istio-init
  kubectl logs -l app=sleep -c istio-init
}

trap dump_logs EXIT

kubectl label ns default istio-injection=enabled --overwrite
istioctl install -y -f ./istio/iop/dual-stack.yaml
kubectl apply -f ./k8s/kind/
kubectl wait --for=condition=ready pod --all --timeout=300s
kubectl get svc whoami -oyaml
kubectl exec -it deploy/sleep -- curl -v whoami:80
