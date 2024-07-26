#!/usr/bin/env bash

set -euo pipefail

# Setup default values
CLUSTER_NAME=${CLUSTER_NAME:-"envoy-gateway"}
METALLB_VERSION=${METALLB_VERSION:-"v0.13.10"}
KIND_NODE_TAG=${KIND_NODE_TAG:-"v1.30.0"}
NUM_WORKERS=${NUM_WORKERS:-""}
RESITRY_MIRROR=${RESITRY_MIRROR:-"192.168.3.73:5000"}
IP_FAMILY=${IP_FAMILY:-"ipv4"}
IP_SPACE=${IPSPACE:-"255"}

if [[ "${IP_FAMILY}" != "ipv4" && "${IP_FAMILY}" != "ipv6" && "${IP_FAMILY}" != "dual" ]]; then
  echo "Invalid IP_FAMILY: ${IP_FAMILY}. Must be ipv4, ipv6 or dual"
  exit 1
fi

  ENVOS=$(uname 2>/dev/null || true)
if [[ "${IP_FAMILY}" == "ipv6" || "${IP_FAMILY}" == "dual" ]]; then
  if [[ "${ENVOS}" != "Linux" ]]; then
    echo "Your system is not supported by this script. Only Linux is supported"
    exit 1
  fi
fi

API_SERVER_ADDRESS=""
if [[ "${ENVOS}" != "Linux" ]]; then
  API_SERVER_ADDRESS="apiServerAddress: 127.0.0.1"
fi

KIND_CFG=$(cat <<-EOM
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  ipFamily: ${IP_FAMILY}
  ${API_SERVER_ADDRESS}
nodes:
- role: control-plane
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors]
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"]
    endpoint = ["http://${RESITRY_MIRROR}", "https://registry-1.docker.io"]
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."quay.io"]
    endpoint = ["http://${RESITRY_MIRROR}", "https://quay.io"]
EOM
)

# https://kind.sigs.k8s.io/docs/user/quick-start/#multi-node-clusters
if [[ -n "${NUM_WORKERS}" ]]; then
for _ in $(seq 1 "${NUM_WORKERS}"); do
  KIND_CFG+=$(printf "\n%s" "- role: worker")
done
fi

## Create kind cluster.
if [[ -z "${KIND_NODE_TAG}" ]]; then
  cat << EOF | kind create cluster --name "${CLUSTER_NAME}" --config -
${KIND_CFG}
EOF
else
  cat << EOF | kind create cluster --image "kindest/node:${KIND_NODE_TAG}" --name "${CLUSTER_NAME}" --config -
${KIND_CFG}
EOF
fi

## Install MetalLB.
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/"${METALLB_VERSION}"/config/manifests/metallb-native.yaml
needCreate="$(kubectl get secret -n metallb-system memberlist --no-headers --ignore-not-found -o custom-columns=NAME:.metadata.name)"
if [ -z "$needCreate" ]; then
    kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey="$(openssl rand -base64 128)"
fi

# Wait for MetalLB to become available.
kubectl rollout status -n metallb-system deployment/controller --timeout 5m
kubectl rollout status -n metallb-system daemonset/speaker --timeout 5m

# Apply config with addresses based on docker network IPAM.
ipv4Prefix=""
ipv6Prefix=""

# Get both ipv4 and ipv6 gateway for the cluster
gatewaystr=$(docker network inspect -f '{{range .IPAM.Config }}{{ .Gateway }} {{end}}' kind | cut -f1,2)
read -r -a gateways <<< "${gatewaystr}"
for gateway in "${gateways[@]}"; do
  if [[ "$gateway" == *"."* ]]; then
    ipv4Prefix=$(echo "${gateway}" |cut -d'.' -f1,2)
  else
    ipv6Prefix=$(echo "${gateway}" |cut -d':' -f1,2,3,4)
  fi
done

ipv4Range="- ${ipv4Prefix}.${IP_SPACE}.200-${ipv4Prefix}.${IP_SPACE}.240"
ipv6Range=""
if [[ "${IP_FAMILY}" == "ipv6" || "${IP_FAMILY}" == "dual" ]]; then
  ipv6Range="- ${ipv6Prefix}::${IP_SPACE}:200-${ipv6Prefix}::${IP_SPACE}:240"
fi

kubectl apply -f - <<EOF
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  namespace: metallb-system
  name: kube-services
spec:
  addresses:
  ${ipv4Range}
  ${ipv6Range}
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


# IPv6 clusters need some CoreDNS changes in order to work in CI:
# Istio CI doesn't offer IPv6 connectivity, so CoreDNS should be configured
# to work in an offline environment:
# https://github.com/coredns/coredns/issues/2494#issuecomment-457215452
# CoreDNS should handle those domains and answer with NXDOMAIN instead of SERVFAIL
# otherwise pods stops trying to resolve the domain.
if [ "${IP_FAMILY}" = "ipv6" ] || [ "${IP_FAMILY}" = "dual" ]; then
  # Get the current config
  original_coredns=$(kubectl get -oyaml -n=kube-system configmap/coredns)
  echo "Original CoreDNS config:"
  echo "${original_coredns}"
  # Patch it
  fixed_coredns=$(
    printf '%s' "${original_coredns}" | sed \
      -e 's/^.*kubernetes cluster\.local/& internal/' \
      -e '/^.*upstream$/d' \
      -e '/^.*fallthrough.*$/d' \
      -e '/^.*forward . \/etc\/resolv.conf$/d' \
      -e '/^.*loop$/d' \
  )
  echo "Patched CoreDNS config:"
  echo "${fixed_coredns}"
  printf '%s' "${fixed_coredns}" | kubectl apply -f -
fi
