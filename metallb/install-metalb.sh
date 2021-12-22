#!/usr/bin/env bash

set -e
set -x



function install_metallb() {
  if [ -z "${METALLB_IPS[*]}" ]; then
    # Take IPs from the end of the docker kind network subnet to use for MetalLB IPs
    DOCKER_KIND_SUBNET="$(docker inspect kind | jq '.[0].IPAM.Config[0].Subnet' -r)"
    METALLB_IPS=()
    while read -r ip; do
      echo $ip ==============
      METALLB_IPS+=("$ip")
    done < <(cidr_to_ips "$DOCKER_KIND_SUBNET" | tail -n 100)
  fi

  # Give this cluster of those IPs
  RANGE="${METALLB_IPS[0]}-${METALLB_IPS[9]}"
  METALLB_IPS=("${METALLB_IPS[@]:10}")

  echo 'apiVersion: v1
kind: ConfigMap
metadata:
  namespace: metallb-system
  name: config
data:
  config: |
    address-pools:
    - name: default
      protocol: layer2
      addresses:
      - '"$RANGE" | kubectl apply -f - 
}


function connect_metallb() {
  REMOTE_NODE=$1
  METALLB_KUBECONFIG=$2
  METALLB_DOCKER_IP=$3

  IP_REGEX='(([0-9]{1,3}\.?){4})'
  LB_CONFIG="$(kubectl --kubeconfig="${METALLB_KUBECONFIG}" -n metallb-system get cm config -o jsonpath="{.data.config}")"
  if [[ "$LB_CONFIG" =~ $IP_REGEX-$IP_REGEX ]]; then
    while read -r lb_cidr; do
      docker exec "${REMOTE_NODE}" ip route add "${lb_cidr}" via "${METALLB_DOCKER_IP}"
    done < <(ips_to_cidrs "${BASH_REMATCH[1]}" "${BASH_REMATCH[3]}")
  fi
}

function cidr_to_ips() {
    CIDR="$1"
    python3 - <<EOF
from ipaddress import IPv4Network; [print(str(ip)) for ip in IPv4Network('$CIDR').hosts()]
EOF
}

function ips_to_cidrs() {
  IP_RANGE_START="$1"
  IP_RANGE_END="$2"
  python3 - <<EOF
from ipaddress import summarize_address_range, IPv4Address
[ print(n.compressed) for n in summarize_address_range(IPv4Address(u'$IP_RANGE_START'), IPv4Address(u'$IP_RANGE_END')) ]
EOF
}
kubectl create ns metallb-system
install_metallb
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.10.2/manifests/metallb.yaml