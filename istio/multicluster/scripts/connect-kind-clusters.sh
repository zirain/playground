#!/usr/bin/env bash

set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

KUBECONFIG_BASE=${KUBECONFIG_BASE:-".kube"}
ISTIO_MC_MODE=${ISTIO_MC_MODE:-"primary-remote"}
ISTIO_NETWORK_MODE=${ISTIO_NETWORK_MODE:-"flat"}

# route_to_cluster SRC DST: let every node of SRC reach the pods and services of DST.
# Each node of DST owns its own pod CIDR, so add one route per DST node, on every SRC node.
function route_to_cluster() {
  local src="${1}" dst="${2}"
  local src_nodes dst_cp_ip dst_svc_cidr
  src_nodes=$(kubectl get nodes --kubeconfig "${KUBECONFIG_BASE}/${src}" -o jsonpath='{.items[*].metadata.name}')
  dst_cp_ip=$(docker inspect -f "{{ .NetworkSettings.Networks.kind.IPAddress }}" "${dst}-control-plane")
  dst_svc_cidr=$(kubectl cluster-info dump --kubeconfig "${KUBECONFIG_BASE}/${dst}" | sed -n 's/^.*--service-cluster-ip-range=\([^"]*\).*$/\1/p' | head -n 1)

  kubectl get nodes --kubeconfig "${KUBECONFIG_BASE}/${dst}" \
    -o jsonpath='{range .items[*]}{.spec.podCIDR} {.status.addresses[?(@.type=="InternalIP")].address}{"\n"}{end}' |
  while read -r pod_cidr node_ip; do
    for node in ${src_nodes}; do
      docker exec "${node}" ip route replace "${pod_cidr}" via "${node_ip}"
    done
  done
  for node in ${src_nodes}; do
    docker exec "${node}" ip route replace "${dst_svc_cidr}" via "${dst_cp_ip}"
  done
}

function connect_kind_clusters() {
  route_to_cluster "${1}" "${2}"
  route_to_cluster "${2}" "${1}"
}

# Clusters on the same Istio network must have direct pod-to-pod connectivity;
# clusters on different networks talk through east-west gateways only.
case "${ISTIO_MC_MODE}/${ISTIO_NETWORK_MODE}" in
  primary-remote/flat)
    # primary, remote1 and remote2 are all on network1
    connect_kind_clusters "remote1" "remote2"
    connect_kind_clusters "primary" "remote1"
    connect_kind_clusters "primary" "remote2"
    ;;
  primary-remote/non-flat)
    # primary and remote1 are on network1, remote2 is on network2
    connect_kind_clusters "primary" "remote1"
    ;;
  *)
    echo "${ISTIO_MC_MODE}/${ISTIO_NETWORK_MODE}: every cluster is on its own network, nothing to connect"
    ;;
esac
