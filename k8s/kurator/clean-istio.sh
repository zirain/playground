#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

export KUBECONFIG=/etc/karmada/karmada-apiserver.config
k delete crd $(k get crd | grep istio.io |awk  '{print $1}')
k delete clusterrole istio-operator
k delete clusterrolebinding istio-operator
k delete ns istio-operator istio-system
k delete cpp istio-customresource-to-primary istio-operator

k delete cop istio-system-member2