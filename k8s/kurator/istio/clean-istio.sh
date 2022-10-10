#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

KUBECTL_FLAGS=' --kubeconfig /etc/karmada/karmada-apiserver.config --ignore-not-found'
CRD_COUNT=$(kubectl get crd --kubeconfig /etc/karmada/karmada-apiserver.config | grep istio.io | wc -l)
if [ ${CRD_COUNT} -gt 0 ]; then
    kubectl delete crd $(kubectl get crd --kubeconfig /etc/karmada/karmada-apiserver.config | grep istio.io |awk  '{print $1}') ${KUBECTL_FLAGS}
fi
kubectl delete clusterrole istio-operator ${KUBECTL_FLAGS}
kubectl delete clusterrolebinding istio-operator ${KUBECTL_FLAGS}
kubectl delete ns istio-operator istio-system ${KUBECTL_FLAGS}
kubectl delete cpp istio-customresource-to-primary istio-operator ${KUBECTL_FLAGS}
kubectl delete cop --all ${KUBECTL_FLAGS}