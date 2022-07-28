#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail


KARMADA_KUBECONFIG=/etc/karmada/karmada-apiserver.config

#kubectl label namespace default istio-injection=enabled --overwrite --kubeconfig=${KARMADA_KUBECONFIG}
#kubectl apply --kubeconfig=${KARMADA_KUBECONFIG} -f https://raw.githubusercontent.com/istio/istio/master/samples/httpbin/sample-client/fortio-deploy.yaml
kubectl apply  -f https://raw.githubusercontent.com/istio/istio/master/samples/sleep/sleep.yaml --kubeconfig=${KARMADA_KUBECONFIG}

kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/helloworld/helloworld.yaml --kubeconfig=${KARMADA_KUBECONFIG}

kubectl apply -f propagation-policy/ --kubeconfig=${KARMADA_KUBECONFIG}