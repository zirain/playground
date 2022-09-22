#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

kubectl delete crd $(k get crd | grep istio.io |awk  '{print $1}') --kubeconfig /etc/karmada/karmada-apiserver.config
kubectl delete clusterrole istio-operator --kubeconfig /etc/karmada/karmada-apiserver.config
kubectl delete clusterrolebinding istio-operator --kubeconfig /etc/karmada/karmada-apiserver.config
kubectl delete ns istio-operator istio-system --kubeconfig /etc/karmada/karmada-apiserver.config
kubectl delete cpp istio-customresource-to-primary istio-operator --kubeconfig /etc/karmada/karmada-apiserver.config
kubectl delete cop --all --kubeconfig /etc/karmada/karmada-apiserver.config