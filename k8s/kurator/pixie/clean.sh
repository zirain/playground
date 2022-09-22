#!/usr/bin/env bash

kubectl delete crd $(kubectl get crd | grep coreos.com |awk  '{print $1}')  --ignore-not-found=true
kubectl delete crd $(kubectl get crd | grep px.dev |awk  '{print $1}')  --ignore-not-found=true
kubectl delete ns px-operator olm px  --ignore-not-found=true
kubectl delete clusterrole aggregate-olm-edit aggregate-olm-view pl-deleter-role system:controller:operator-lifecycle-manager pl-deleter-binding operator-lifecycle-manager  --ignore-not-found=true
kubectl delete clusterrolebinding olm-operator-cluster-binding-olm pl-deleter-binding  --ignore-not-found=true
kubectl delete sa pl-deleter-service-account  --ignore-not-found=true
kubectl delete cpp --all --ignore-not-found=true
