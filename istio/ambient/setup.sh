#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

#set -x

# This script starts a local karmada control plane based on current codebase and with a certain number of clusters joined.
# Parameters: [HOST_IPADDRESS](optional) if you want to export clusters' API server port to specific IP address
# This script depends on utils in: ${REPO_ROOT}/hack/util.sh
# 1. used by developer to setup develop environment quickly.
# 2. used by e2e testing to setup test environment automatically.
REPO_ROOT=$(dirname "${BASH_SOURCE[0]}")
source "${REPO_ROOT}"/util.sh

# variable define
KUBECONFIG_PATH=${KUBECONFIG_PATH:-"${HOME}/.kube"}
AMBIENT_KUBECONFIG=${AMBIENT_KUBECONFIG:-"${KUBECONFIG_PATH}/istio-ambient.config"}
CLUSTER_NAME=${CLUSTER_NAME:-"ambient"}

HOST_IPADDRESS=${1:-}

METALLB_VERSION=${METALLB_VERSION:-"v0.10.2"}
CLUSTER_VERSION=${CLUSTER_VERSION:-"kindest/node:v1.24.0"}
KIND_LOG_FILE=${KIND_LOG_FILE:-"/tmp/istio"}
AMBIENT_IMAGE_TAG=${AMBIENT_IMAGE_TAG:-"0.0.0-ambient.191fe680b52c1754ee72a06b3e0d3f9d116f2e82"}
BOOKINFO_IMAGE_TAG=${BOOKINFO_IMAGE_TAG:-"1.17.0"}

# create host cluster and member clusters in parallel
# host IP address: script parameter ahead of macOS IP

#prepare for kindClusterConfig
TEMP_PATH=$(mktemp -d)
echo -e "Preparing kindClusterConfig in path: ${TEMP_PATH}"
cp -rf ambient.yaml "${TEMP_PATH}"/ambient.yaml

util::create_cluster "${CLUSTER_NAME}" "${AMBIENT_KUBECONFIG}" "${CLUSTER_VERSION}" "${KIND_LOG_FILE}" "${TEMP_PATH}"/ambient.yaml

util::check_clusters_ready "${AMBIENT_KUBECONFIG}" "${CLUSTER_NAME}"

IMAGES=("quay.io/metallb/controller:${METALLB_VERSION}" \
"quay.io/metallb/speaker:${METALLB_VERSION}" \
"gcr.io/istio-testing/install-cni:${AMBIENT_IMAGE_TAG}" \
"gcr.io/istio-testing/proxyv2:${AMBIENT_IMAGE_TAG}" \
"gcr.io/istio-testing/pilot:${AMBIENT_IMAGE_TAG}" \
"istio/examples-bookinfo-details-v1:${BOOKINFO_IMAGE_TAG}" \
"istio/examples-bookinfo-productpage-v1:${BOOKINFO_IMAGE_TAG}" \
"istio/examples-bookinfo-ratings-v1:${BOOKINFO_IMAGE_TAG}" \
"istio/examples-bookinfo-reviews-v1:${BOOKINFO_IMAGE_TAG}" \
"istio/examples-bookinfo-reviews-v2:${BOOKINFO_IMAGE_TAG}" \
"istio/examples-bookinfo-reviews-v3:${BOOKINFO_IMAGE_TAG}" \
"curlimages/curl")
for img in "${IMAGES[@]}"; do
    docker pull ${img}
done

KIND_CLUSTES=("${CLUSTER_NAME}")
for cluster in "${KIND_CLUSTES[@]}"; do
    for img in "${IMAGES[@]}"; do
        echo "load image ${img} to ${cluster}"
        kind load docker-image ${img} --name ${cluster}
    done

    echo "install metallb in $cluster"
    kubectl create ns metallb-system --kubeconfig="${AMBIENT_KUBECONFIG}" --context="${cluster}"
    util::install_metallb ${AMBIENT_KUBECONFIG} ${cluster}
    kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/${METALLB_VERSION}/manifests/metallb.yaml --kubeconfig="${AMBIENT_KUBECONFIG}" --context="${cluster}"
done

echo "install istio in ambient cluster"

istioctl-ambient install --kubeconfig=${AMBIENT_KUBECONFIG} --context="${CLUSTER_NAME}" --set profile=ambient -y --set meshConfig.accessLogFile=/dev/stdout

kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/bookinfo/platform/kube/bookinfo.yaml --kubeconfig=${AMBIENT_KUBECONFIG} --context="${CLUSTER_NAME}"
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/bookinfo/networking/bookinfo-gateway.yaml --kubeconfig=${AMBIENT_KUBECONFIG} --context="${CLUSTER_NAME}"

kubectl apply -f https://raw.githubusercontent.com/linsun/sample-apps/main/sleep/sleep.yaml --kubeconfig=${AMBIENT_KUBECONFIG} --context="${CLUSTER_NAME}"
kubectl apply -f https://raw.githubusercontent.com/linsun/sample-apps/main/sleep/notsleep.yaml --kubeconfig=${AMBIENT_KUBECONFIG} --context="${CLUSTER_NAME}"

# Adding your application to the ambient mesh
kubectl label namespace default istio.io/dataplane-mode=ambient --kubeconfig=${AMBIENT_KUBECONFIG} --context="${CLUSTER_NAME}"


kubectl apply -f networking/productpage-gateway.yaml --kubeconfig=${AMBIENT_KUBECONFIG} --context="${CLUSTER_NAME}"
kubectl apply -f networking/reviews-gateway.yaml --kubeconfig=${AMBIENT_KUBECONFIG} --context="${CLUSTER_NAME}"
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/bookinfo/networking/virtual-service-reviews-90-10.yaml --kubeconfig=${AMBIENT_KUBECONFIG} --context="${CLUSTER_NAME}"
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/bookinfo/networking/destination-rule-reviews.yaml --kubeconfig=${AMBIENT_KUBECONFIG} --context="${CLUSTER_NAME}"

function print_success() {
  echo "Multicluster Istio is running."
  echo -e "\nTo manage your clusters, run:"
  echo -e "  export KUBECONFIG=${AMBIENT_KUBECONFIG}"
  echo "Please use 'kubectl config use-context ambient' to switch to the different remote cluster."
}

print_success
