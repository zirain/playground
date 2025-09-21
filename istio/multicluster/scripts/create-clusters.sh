#!/usr/bin/env bash

set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

SCRIPT_DIR=$(dirname $(realpath "$0"))
BASE_DIR="${SCRIPT_DIR}/.."

IP_FAMILY="${IP_FAMILY:-"ipv4"}"
KIND_NODE_TAG="${KIND_NODE_TAG:-"v1.33.0"}"
ISTIO_MC_MODE=${ISTIO_MC_MODE:-"primary-remote"}
ISTIO_NETWORK_MODE=${ISTIO_NETWORK_MODE:-"flat"}
OS="$(uname)"

KIND_CFG_PREFIX="${BASE_DIR}/${ISTIO_MC_MODE}/kind-configs"
KIND_CFG_FILES=$(ls ${KIND_CFG_PREFIX})

KUBECONFIG_BASE=${KUBECONFIG_BASE:-".kube"}
mkdir -p "${KUBECONFIG_BASE}"

function check_cluster_ready(){
  CLUSTER_NAME=$1
  CONTAINER_IP=$(docker inspect "${CLUSTER_NAME}-control-plane" --format "{{ .NetworkSettings.Networks.kind.IPAddress }}")
  (kind get kubeconfig --name "${CLUSTER_NAME}" --internal | sed "s/${CLUSTER_NAME}-control-plane/${CONTAINER_IP}/g") > "${KUBECONFIG_BASE}/${CLUSTER_NAME}"
}

ipv4_address_start_at=100
for file in ${KIND_CFG_FILES[@]}; do
  file_name=$(basename ${file})
  file_full_path="${KIND_CFG_PREFIX}/${file}"
  cluster_name="${file_name%.*}"
  kind create cluster --name "${cluster_name}" --image "kindest/node:${KIND_NODE_TAG}" --config "${file_full_path}" --kubeconfig "${KUBECONFIG_BASE}/${cluster_name}"
  # Wait for kind cluster ready
  while true; do
    if check_cluster_ready "${cluster_name}"; then
      break
    fi
    sleep 3s
  done

  # Apply config with addresses based on docker network IPAM.
  address_ranges=""
  if [ "${IP_FAMILY}" = "ipv4" ] || [ "${IP_FAMILY}" = "dual" ]; then
      subnet_v4=$(docker network inspect kind | jq -r '.[].IPAM.Config[] | select(.Subnet | contains(":") | not) | .Subnet')
      address_prefix_v4=$(echo "${subnet_v4}" | awk -F. '{print $1"."$2"."$3}')
      address_range_v4="${address_prefix_v4}.${ipv4_address_start_at}-${address_prefix_v4}.250"
      echo "IPv4 address range: ${address_range_v4}"
      address_ranges+="- ${address_range_v4}"
  fi
  # plus 50 to ipv4_address_start_at
  ipv4_address_start_at=$((ipv4_address_start_at + 50))

  if [ "${IP_FAMILY}" = "ipv6" ] || [ "${IP_FAMILY}" = "dual" ]; then
      subnet_v6=$(docker network inspect kind | jq -r '.[].IPAM.Config[] | select(.Subnet | contains(":")) | .Subnet')
      ipv6_prefix="${subnet_v6%::*}"
      address_range_v6="${ipv6_prefix}::c8-${ipv6_prefix}::fa"
      echo "IPv6 address range: ${address_range_v6}"
      [ -n "${address_ranges}" ] && address_ranges+="\n"
      address_ranges+="- ${address_range_v6}"
  fi

  echo "Starting metallb deployment in cluster ${cluster_name}"
  kubectl apply -f "${BASE_DIR}/addons/metallb-native.yaml" --kubeconfig "${KUBECONFIG_BASE}/${CLUSTER_NAME}"
  kubectl wait --for=condition=available --timeout=5m deployment/controller -n metallb-system --kubeconfig "${KUBECONFIG_BASE}/${CLUSTER_NAME}"
  kubectl apply --kubeconfig "${KUBECONFIG_BASE}/${CLUSTER_NAME}" -f - <<EOF >/dev/null 2>&1
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  namespace: metallb-system
  name: kube-services
spec:
  addresses:
$(echo -e "${address_ranges}" | sed 's/^/    /')
---
apiVersion: metallb.io/v1beta1
kind: L2Advertisement
metadata:
  name: kube-services
  namespace: metallb-system
spec:
  ipAddressPools:
  - kube-services
EOF
done

# echo "Kind CIDR is $(docker network inspect -f '{{$map := index .IPAM.Config 0}}{{index $map "Subnet"}}' kind)"
echo "Complete"
echo "You can find the kubeconfig files in ${KUBECONFIG_BASE}"