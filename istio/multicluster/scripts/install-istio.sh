#!/usr/bin/env bash

set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

SCRIPT_DIR=$(dirname $(realpath "$0"))
BASE_DIR="${SCRIPT_DIR}/.."

KUBECONFIG_BASE=${KUBECONFIG_BASE:-".kube"}

ISTIO_MC_MODE=${ISTIO_MC_MODE:-"primary-remote"}
ISTIO_NETWORK_MODE=${ISTIO_NETWORK_MODE:-"flat"}
ISTIO_BASE_DIR="${BASE_DIR}/${ISTIO_MC_MODE}/${ISTIO_NETWORK_MODE}"
IOP_CFG_PREFIX="${ISTIO_BASE_DIR}/iop"
IOP_CFG_FILES=$(ls ${IOP_CFG_PREFIX})

function install_primary_remote() {
    echo "Install Istio on primary"
    istioctl install -y -f "${IOP_CFG_PREFIX}/primary.yaml"  --kubeconfig "${KUBECONFIG_BASE}/primary"
    kubectl apply -f "${ISTIO_BASE_DIR}/expose-istiod.yaml" -n istio-system --kubeconfig "${KUBECONFIG_BASE}/primary"
    DISCOVER_ADDRESS=$(kubectl get --kubeconfig "${KUBECONFIG_BASE}/primary" svc -nistio-system istio-eastwestgateway  -o jsonpath="{.status.loadBalancer.ingress[0].ip}")

    echo "Install Istio on remote1"
    istioctl create-remote-secret --kubeconfig "${KUBECONFIG_BASE}/remote1" --name=remote1 | kubectl apply -f - --kubeconfig "${KUBECONFIG_BASE}/primary"
    # sync root ca
    kubectl --kubeconfig="${KUBECONFIG_BASE}/remote1" create namespace istio-system || true
    kubectl --kubeconfig="${KUBECONFIG_BASE}/remote1" annotate namespace istio-system topology.istio.io/controlPlaneClusters=primary
    istioctl install -y -f "${IOP_CFG_PREFIX}/remote1.yaml" --kubeconfig "${KUBECONFIG_BASE}/remote1" \
        --set values.global.remotePilotAddress="${DISCOVER_ADDRESS}" \
        --set profile=remote

    echo "Install Istio on remote2"
    istioctl create-remote-secret --kubeconfig "${KUBECONFIG_BASE}/remote2" --name=remote2 | kubectl apply -f - --kubeconfig "${KUBECONFIG_BASE}/primary"
    # sync root ca
    kubectl --kubeconfig="${KUBECONFIG_BASE}/remote2" create namespace istio-system || true
    kubectl --kubeconfig="${KUBECONFIG_BASE}/remote2" annotate namespace istio-system topology.istio.io/controlPlaneClusters=primary
    istioctl install -y -f "${IOP_CFG_PREFIX}/remote2.yaml" --kubeconfig "${KUBECONFIG_BASE}/remote2" \
        --set values.global.remotePilotAddress="${DISCOVER_ADDRESS}" \
        --set profile=remote
    if [ "${ISTIO_NETWORK_MODE}" = "non-flat" ]; then
      echo "Install east-west gateway on remote2"
      istioctl install -y -f "${IOP_CFG_PREFIX}/remote2-eastwest.yaml" --kubeconfig "${KUBECONFIG_BASE}/remote2"
    fi
}

install_primary_remote