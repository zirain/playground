#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

#  set -x

# This script starts a local karmada control plane based on current codebase and with a certain number of clusters joined.
# Parameters: [HOST_IPADDRESS](optional) if you want to export clusters' API server port to specific IP address
# This script depends on utils in: ${REPO_ROOT}/hack/util.sh
# 1. used by developer to setup develop environment quickly.
# 2. used by e2e testing to setup test environment automatically.
REPO_ROOT=$(dirname "${BASH_SOURCE[0]}")
source "${REPO_ROOT}"/util.sh

# variable define
WITH_PROMETHEUS=${WITH_PROMETHEUS:-'false'}
KUBECONFIG_PATH=${KUBECONFIG_PATH:-"${HOME}/.kube"}
MAIN_KUBECONFIG=${MAIN_KUBECONFIG:-"${KUBECONFIG_PATH}/istio-primary.config"}
MEMBER_CLUSTER_KUBECONFIG=${MEMBER_CLUSTER_KUBECONFIG:-"${KUBECONFIG_PATH}/istio-remote.config"}
PRIMARY_CLUSTER_NAME=${PRIMARY_CLUSTER_NAME:-"primary"}
REMOTE_CLUSTER_NAME=${REMOTE_CLUSTER_NAME:-"remote"}
HOST_IPADDRESS=${1:-}

METALLB_VERSION=${METALLB_VERSION:-"v0.13.10"}
CLUSTER_VERSION=${CLUSTER_VERSION:-"kindest/node:v1.28.0"}
KIND_LOG_FILE=${KIND_LOG_FILE:-"/tmp/istio"}

# create host cluster and member clusters in parallel
# host IP address: script parameter ahead of macOS IP

#prepare for kindClusterConfig
TEMP_PATH=$(mktemp -d)
KIND_CONFIGS_PATH="./kind-configs"
echo -e "Preparing kindClusterConfig in path: ${TEMP_PATH}"
cp -rf ${KIND_CONFIGS_PATH}/primary.yaml "${TEMP_PATH}"/primary.yaml
cp -rf ${KIND_CONFIGS_PATH}/remote.yaml "${TEMP_PATH}"/remote.yaml

util::create_cluster "${PRIMARY_CLUSTER_NAME}" "${MAIN_KUBECONFIG}" "${CLUSTER_VERSION}" "${KIND_LOG_FILE}" "${TEMP_PATH}"/primary.yaml
util::create_cluster "${REMOTE_CLUSTER_NAME}" "${MEMBER_CLUSTER_KUBECONFIG}" "${CLUSTER_VERSION}" "${KIND_LOG_FILE}" "${TEMP_PATH}"/remote.yaml

util::check_clusters_ready "${MAIN_KUBECONFIG}" "${PRIMARY_CLUSTER_NAME}"
sleep 5s
util::check_clusters_ready "${MEMBER_CLUSTER_KUBECONFIG}" "${REMOTE_CLUSTER_NAME}"
sleep 5s

# load images
docker pull quay.io/metallb/controller:${METALLB_VERSION}
docker pull quay.io/metallb/speaker:${METALLB_VERSION}

KIND_CLUSTES=("${PRIMARY_CLUSTER_NAME}" "${REMOTE_CLUSTER_NAME}")
for cluster in "${KIND_CLUSTES[@]}"; do
    echo "load image to ${cluster}"
    kind load docker-image quay.io/metallb/controller:${METALLB_VERSION} --name ${cluster}
    kind load docker-image quay.io/metallb/speaker:${METALLB_VERSION} --name ${cluster}
done

echo "connect primary <-> remote1"
util::connect_kind_clusters "${PRIMARY_CLUSTER_NAME}" "${MAIN_KUBECONFIG}" "${REMOTE_CLUSTER_NAME}" "${MEMBER_CLUSTER_KUBECONFIG}" 1

echo "cluster networks connected"

echo "starting install metallb in primary cluster"
kubectl create ns metallb-system --kubeconfig="${MAIN_KUBECONFIG}" --context="${PRIMARY_CLUSTER_NAME}"
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/"${METALLB_VERSION}"/config/manifests/metallb-native.yaml --kubeconfig="${MAIN_KUBECONFIG}" --context="${PRIMARY_CLUSTER_NAME}"
kubectl rollout status --watch --timeout=5m -n metallb-system deploy/controller --kubeconfig="${MAIN_KUBECONFIG}" --context="${PRIMARY_CLUSTER_NAME}"
util::install_metallb ${MAIN_KUBECONFIG} ${PRIMARY_CLUSTER_NAME}

echo "starting install metallb in member clusters"
MEMBER_CLUSTERS=(${REMOTE_CLUSTER_NAME})
for c in ${MEMBER_CLUSTERS[@]}
do
  echo "install metallb in $c"
  kubectl create ns metallb-system --kubeconfig="${MEMBER_CLUSTER_KUBECONFIG}" --context="${c}"
  kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/"${METALLB_VERSION}"/config/manifests/metallb-native.yaml --kubeconfig="${MEMBER_CLUSTER_KUBECONFIG}" --context="${c}"
  kubectl rollout status --watch --timeout=5m -n metallb-system deploy/controller --kubeconfig="${MEMBER_CLUSTER_KUBECONFIG}" --context="${c}"
  util::install_metallb ${MEMBER_CLUSTER_KUBECONFIG} ${c}
done

function print_success() {
  echo "Multicluster is running."
  echo -e "\nTo start using your Istio, run:"
  echo -e "  export KUBECONFIG=${MAIN_KUBECONFIG}"
  echo -e "\nTo manage your remote clusters, run:"
  echo -e "  export KUBECONFIG=${MEMBER_CLUSTER_KUBECONFIG}"
  echo "Please use 'kubectl config use-context remote' to switch to the different remote cluster."
}

print_success
