#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail


export KUBECONFIG=/etc/karmada/karmada-apiserver.config
k delete crd $(k get crd | grep coreos.com |awk  '{print $1}')
k delete clusterrole prometheus-operator prometheus kube-state-metrics prometheus-k8s
k delete clusterrolebinding prometheus-operator prometheus kube-state-metrics prometheus-k8s
k delete ns monitoring
k delete cpp prometheus-operator prometheus-rbac
k delete ns prometheus
k delete role prometheus-k8s -n default
k delete role prometheus-k8s -n kube-system
k delete rolebinding prometheus-k8s -n default
k delete rolebinding prometheus-k8s -n kube-system

