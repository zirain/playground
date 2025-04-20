#!/usr/bin/env bash

set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

KUBECONFIG_BASE=${KUBECONFIG_BASE:-".kube"}

function connect_kind_clusters() {
  C1="${1}"
  C2="${2}"
  POD_TO_POD_AND_SERVICE_CONNECTIVITY="${3}"

  C1_NODE="${C1}-control-plane"
  C2_NODE="${C2}-control-plane"
  C1_DOCKER_IP=$(docker inspect -f "{{ .NetworkSettings.Networks.kind.IPAddress }}" "${C1_NODE}")
  C2_DOCKER_IP=$(docker inspect -f "{{ .NetworkSettings.Networks.kind.IPAddress }}" "${C2_NODE}")
  if [ "${POD_TO_POD_AND_SERVICE_CONNECTIVITY}" -eq 1 ]; then
    # Set up routing rules for inter-cluster direct pod to pod & service communication
    C1_POD_CIDR=$(kubectl get node --kubeconfig "${KUBECONFIG_BASE}/${C1}" -ojsonpath='{.items[0].spec.podCIDR}')
    C2_POD_CIDR=$(kubectl get node --kubeconfig "${KUBECONFIG_BASE}/${C2}" -ojsonpath='{.items[0].spec.podCIDR}')
    C1_SVC_CIDR=$(kubectl cluster-info dump --kubeconfig "${KUBECONFIG_BASE}/${C1}" | sed -n 's/^.*--service-cluster-ip-range=\([^"]*\).*$/\1/p' | head -n 1)
    C2_SVC_CIDR=$(kubectl cluster-info dump --kubeconfig "${KUBECONFIG_BASE}/${C2}" | sed -n 's/^.*--service-cluster-ip-range=\([^"]*\).*$/\1/p' | head -n 1)
    docker exec "${C1_NODE}" ip route add "${C2_POD_CIDR}" via "${C2_DOCKER_IP}"
    docker exec "${C1_NODE}" ip route add "${C2_SVC_CIDR}" via "${C2_DOCKER_IP}"
    docker exec "${C2_NODE}" ip route add "${C1_POD_CIDR}" via "${C1_DOCKER_IP}"
    docker exec "${C2_NODE}" ip route add "${C1_SVC_CIDR}" via "${C1_DOCKER_IP}"
  fi
}

ISTIO_NETWORK_MODE=${ISTIO_NETWORK_MODE:-"flat"}
if [ "${ISTIO_NETWORK_MODE}" = "flat" ]; then
    echo "connect clusters in flat mode"

    # connecting networks between primary, remote1 and remote2 clusters
    echo "connect remote1 <-> remote2"
    connect_kind_clusters "remote1" "remote2" 1

    echo "connect primary <-> remote1"
    connect_kind_clusters "primary" "remote1" 1

    echo "connect primary <-> remote2"
    connect_kind_clusters "primary" "remote2" 1
fi