#!/usr/bin/env bash

set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

SCRIPT_DIR=$(dirname $(realpath "$0"))
BASE_DIR="${SCRIPT_DIR}/.."

KUBECONFIG_BASE=${KUBECONFIG_BASE:-".kube"}

# Install sleep in primary
kubectl apply -f "${BASE_DIR}/addons/sidecar.yaml" --kubeconfig "${KUBECONFIG_BASE}/primary"
kubectl apply -n sidecar -f "${BASE_DIR}/addons/helloworld/helloworld-svc.yaml" --kubeconfig "${KUBECONFIG_BASE}/primary"
kubectl apply -n sidecar -f "${BASE_DIR}/addons/sleep/sleep.yaml" --kubeconfig "${KUBECONFIG_BASE}/primary"
kubectl apply -n sidecar -f "${BASE_DIR}/addons/helloworld/helloworld-v1.yaml" --kubeconfig "${KUBECONFIG_BASE}/primary"
kubectl apply -n sidecar -f "${BASE_DIR}/addons/helloworld/helloworld-v2.yaml" --kubeconfig "${KUBECONFIG_BASE}/primary"

echo "Install helloworld-v1 in remote1"
kubectl apply -f "${BASE_DIR}/addons/sidecar.yaml" --kubeconfig "${KUBECONFIG_BASE}/remote1"
kubectl apply -n sidecar -f "${BASE_DIR}/addons/helloworld/helloworld-v1.yaml" --kubeconfig "${KUBECONFIG_BASE}/remote1"
kubectl apply -n sidecar -f "${BASE_DIR}/addons/helloworld/helloworld-v2.yaml" --kubeconfig "${KUBECONFIG_BASE}/remote1"
    
echo "Install helloworld-v2 in remote2"
kubectl apply -f "${BASE_DIR}/addons/sidecar.yaml" --kubeconfig "${KUBECONFIG_BASE}/remote2"
kubectl apply -n sidecar -f "${BASE_DIR}/addons/helloworld/helloworld-v1.yaml" --kubeconfig "${KUBECONFIG_BASE}/remote2"
kubectl apply -n sidecar -f "${BASE_DIR}/addons/helloworld/helloworld-v2.yaml" --kubeconfig "${KUBECONFIG_BASE}/remote2"
