#!/usr/bin/env bash

set -e
set -u
set -o pipefail
set -x

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

VM_APP=${VM_APP:-"vm-app"}
VM_NAMESPACE=${VM_NAMESPACE:-"vm-namespace"}
WORK_DIR=${WORK_DIR:-"$(mktemp -d)"}
SERVICE_ACCOUNT=${SERVICE_ACCOUNT:-"default"}
CLUSTER_NETWORK=${CLUSTER_NETWORK:-"network1"}
VM_NETWORK=${VM_NETWORK:-"network1"}
CLUSTER=${CLUSTER:-"Kubernetes"}


# Install Istio
istioctl install -f "${SCRIPT_DIR}/iop.yaml" -y
# Install East-West Gateway
istioctl install -y -f "${SCRIPT_DIR}/eastwest.yaml"
# Expose Istiod service
kubectl apply -f "${SCRIPT_DIR}/expose-istiod.yaml"
# Create VM namespace
kubectl create namespace "${VM_NAMESPACE}" || true
kubectl create serviceaccount "${SERVICE_ACCOUNT}" -n "${VM_NAMESPACE}" || true


WORKLOAD_DIR="${SCRIPT_DIR}/out"

istioctl x workload entry configure -f ${SCRIPT_DIR}/workloadgroup.yaml -o "${WORKLOAD_DIR}/" --clusterID "${CLUSTER}"

sudo mkdir -p /etc/certs
sudo cp "${WORKLOAD_DIR}"/root-cert.pem /etc/certs/root-cert.pem

sudo  mkdir -p /var/run/secrets/tokens
sudo cp "${WORKLOAD_DIR}"/istio-token /var/run/secrets/tokens/istio-token



# curl -LO https://storage.googleapis.com/istio-release/releases/1.24.2/deb/istio-sidecar.deb
sudo dpkg -i "${SCRIPT_DIR}/istio-sidecar-1.24.2.deb"
sudo cp "${WORKLOAD_DIR}"/cluster.env /var/lib/istio/envoy/cluster.env
sudo cp "${WORKLOAD_DIR}"/mesh.yaml /etc/istio/config/mesh
# sudo sh -c 'cat ${WORKLOAD_DIR}/hosts >> /etc/hosts'

sudo mkdir -p /etc/istio/proxy
sudo chown -R istio-proxy /var/lib/istio /etc/certs /etc/istio/proxy /etc/istio/config /var/run/secrets /etc/certs/root-cert.pem

sudo systemctl start istio

