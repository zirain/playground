#!/usr/bin/env bash

set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

SCRIPT_DIR=$(dirname $(realpath "$0"))
BASE_DIR="${SCRIPT_DIR}/.."

KUBECONFIG_BASE=${KUBECONFIG_BASE:-".kube"}
ISTIO_MC_MODE=${ISTIO_MC_MODE:-"primary-remote"}

function install_ambient_multi_primary_app() {
    for i in 1 2; do
        kubeconfig="${KUBECONFIG_BASE}/cluster${i}"
        echo "Install sleep and helloworld-v${i} in cluster${i}"
        kubectl apply -f "${BASE_DIR}/addons/ambient.yaml" --kubeconfig "${kubeconfig}"
        kubectl apply -n ambient -f "${BASE_DIR}/addons/helloworld/helloworld-svc.yaml" --kubeconfig "${kubeconfig}"
        kubectl apply -n ambient -f "${BASE_DIR}/addons/sleep/sleep.yaml" --kubeconfig "${kubeconfig}"
        kubectl apply -n ambient -f "${BASE_DIR}/addons/helloworld/helloworld-v${i}.yaml" --kubeconfig "${kubeconfig}"
        # ambient only shares services across clusters when they are marked global
        kubectl label svc helloworld -n ambient istio.io/global=true --overwrite --kubeconfig "${kubeconfig}"
    done
}

if [ "${ISTIO_MC_MODE}" = "ambient-multi-primary" ]; then
    install_ambient_multi_primary_app
    exit 0
fi

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
