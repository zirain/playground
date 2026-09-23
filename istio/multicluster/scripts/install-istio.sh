#!/usr/bin/env bash

set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

SCRIPT_DIR=$(dirname $(realpath "$0"))
BASE_DIR="${SCRIPT_DIR}/.."

KUBECONFIG_BASE=${KUBECONFIG_BASE:-".kube"}
CERTS_DIR=${CERTS_DIR:-".certs"}

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
    kubectl --kubeconfig="${KUBECONFIG_BASE}/remote1" create namespace istio-system || true
    kubectl --kubeconfig="${KUBECONFIG_BASE}/remote1" annotate namespace istio-system topology.istio.io/controlPlaneClusters=primary
    kubectl --kubeconfig="${KUBECONFIG_BASE}/remote1" label namespace istio-system topology.istio.io/network=network1 --overwrite
    istioctl install -y -f "${IOP_CFG_PREFIX}/remote1.yaml" --kubeconfig "${KUBECONFIG_BASE}/remote1" \
        --set values.global.remotePilotAddress="${DISCOVER_ADDRESS}" \
        --set profile=remote
    # attach remote1 as a remote cluster of primary
    istioctl create-remote-secret --kubeconfig "${KUBECONFIG_BASE}/remote1" --name=remote1 | kubectl apply -f - --kubeconfig "${KUBECONFIG_BASE}/primary"

    echo "Install Istio on remote2"
    kubectl --kubeconfig="${KUBECONFIG_BASE}/remote2" create namespace istio-system || true
    kubectl --kubeconfig="${KUBECONFIG_BASE}/remote2" annotate namespace istio-system topology.istio.io/controlPlaneClusters=primary
    kubectl --kubeconfig="${KUBECONFIG_BASE}/remote2" label namespace istio-system topology.istio.io/network=network1 --overwrite
    if [ "${ISTIO_NETWORK_MODE}" = "non-flat" ]; then
        kubectl --kubeconfig="${KUBECONFIG_BASE}/remote2" label namespace istio-system topology.istio.io/network=network2 --overwrite
    fi
    istioctl install -y -f "${IOP_CFG_PREFIX}/remote2.yaml" --kubeconfig "${KUBECONFIG_BASE}/remote2" \
        --set values.global.remotePilotAddress="${DISCOVER_ADDRESS}" \
        --set profile=remote

    # attach remote2 as a remote cluster of primary
    istioctl create-remote-secret --kubeconfig "${KUBECONFIG_BASE}/remote2" --name=remote2 | kubectl apply -f - --kubeconfig "${KUBECONFIG_BASE}/primary"
    if [ "${ISTIO_NETWORK_MODE}" = "non-flat" ]; then
      echo "Install east-west gateway on remote2"
      istioctl install -y -f "${IOP_CFG_PREFIX}/remote2-eastwest.yaml" --kubeconfig "${KUBECONFIG_BASE}/remote2"
    fi
}

function install_ambient_multi_primary() {
    CLUSTERS=(cluster1 cluster2)
    GATEWAY_API_VERSION=${GATEWAY_API_VERSION:-"v1.4.0"}

    # multi-primary requires all istiods to share a common root of trust
    "${SCRIPT_DIR}/gen-certs.sh" "${CLUSTERS[@]}"

    for i in "${!CLUSTERS[@]}"; do
        cluster="${CLUSTERS[$i]}"
        network="network$((i + 1))"
        kubeconfig="${KUBECONFIG_BASE}/${cluster}"

        echo "Install Istio ambient on ${cluster} (${network})"
        kubectl get crd gateways.gateway.networking.k8s.io --kubeconfig "${kubeconfig}" &> /dev/null || \
            kubectl apply --server-side --kubeconfig "${kubeconfig}" \
                -f "https://github.com/kubernetes-sigs/gateway-api/releases/download/${GATEWAY_API_VERSION}/experimental-install.yaml"
        kubectl --kubeconfig="${kubeconfig}" create namespace istio-system || true
        kubectl --kubeconfig="${kubeconfig}" label namespace istio-system topology.istio.io/network="${network}" --overwrite
        kubectl --kubeconfig="${kubeconfig}" create secret generic cacerts -n istio-system \
            --from-file="${CERTS_DIR}/${cluster}/ca-cert.pem" \
            --from-file="${CERTS_DIR}/${cluster}/ca-key.pem" \
            --from-file="${CERTS_DIR}/${cluster}/root-cert.pem" \
            --from-file="${CERTS_DIR}/${cluster}/cert-chain.pem" \
            --dry-run=client -o yaml | kubectl apply -f - --kubeconfig "${kubeconfig}"
        istioctl install -y -f "${IOP_CFG_PREFIX}/${cluster}.yaml" --kubeconfig "${kubeconfig}"

        echo "Install east-west gateway on ${cluster}"
        kubectl apply -f "${ISTIO_BASE_DIR}/eastwest/${cluster}.yaml" --kubeconfig "${kubeconfig}"
        kubectl wait --for=condition=programmed --timeout=5m gateway/istio-eastwestgateway -n istio-system --kubeconfig "${kubeconfig}"
    done

    # every primary watches the API servers of all other clusters for endpoint discovery
    for src in "${CLUSTERS[@]}"; do
        for dst in "${CLUSTERS[@]}"; do
            [ "${src}" = "${dst}" ] && continue
            istioctl create-remote-secret --kubeconfig "${KUBECONFIG_BASE}/${src}" --name="${src}" | \
                kubectl apply -f - --kubeconfig "${KUBECONFIG_BASE}/${dst}"
        done
    done
}

case "${ISTIO_MC_MODE}" in
    primary-remote) install_primary_remote ;;
    ambient-multi-primary) install_ambient_multi_primary ;;
    *) echo "unsupported ISTIO_MC_MODE: ${ISTIO_MC_MODE}"; exit 1 ;;
esac
