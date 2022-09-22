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
MEMBER_CLUSTER_KUBECONFIG=${MEMBER_CLUSTER_KUBECONFIG:-"${KUBECONFIG_PATH}/istio-remotes.config"}
CLUSTER1_NAME=${CLUSTER1_NAME:-"cluster1"}
CLUSTER2_NAME=${CLUSTER2_NAME:-"cluster2"}
HOST_IPADDRESS=${1:-}

ISTIO_TAG=${ISTIO_TAG:-"1.14.3"}
METALLB_VERSION=${METALLB_VERSION:-"v0.10.2"}
CLUSTER_VERSION=${CLUSTER_VERSION:-"kindest/node:v1.24.0"}
KIND_LOG_FILE=${KIND_LOG_FILE:-"/tmp/istio"}

# create host cluster and member clusters in parallel
# host IP address: script parameter ahead of macOS IP

#prepare for kindClusterConfig
TEMP_PATH=$(mktemp -d)
echo -e "Preparing kindClusterConfig in path: ${TEMP_PATH}"
cp -rf cluster1.yaml "${TEMP_PATH}"/cluster1.yaml
cp -rf cluster2.yaml "${TEMP_PATH}"/cluster2.yaml

util::create_cluster "${CLUSTER1_NAME}" "${MEMBER_CLUSTER_KUBECONFIG}" "${CLUSTER_VERSION}" "${KIND_LOG_FILE}" "${TEMP_PATH}"/cluster1.yaml
util::create_cluster "${CLUSTER2_NAME}" "${MEMBER_CLUSTER_KUBECONFIG}" "${CLUSTER_VERSION}" "${KIND_LOG_FILE}" "${TEMP_PATH}"/cluster2.yaml

sleep 30s
util::check_clusters_ready "${MEMBER_CLUSTER_KUBECONFIG}" "${CLUSTER1_NAME}"
sleep 5s
util::check_clusters_ready "${MEMBER_CLUSTER_KUBECONFIG}" "${CLUSTER2_NAME}"
sleep 5s

IMAGES=("quay.io/metallb/controller:${METALLB_VERSION}" "quay.io/metallb/speaker:${METALLB_VERSION}" "istio/proxyv2:${ISTIO_TAG}" "istio/pilot:${ISTIO_TAG}")
for img in "${IMAGES[@]}"; do
    docker pull ${img}
done

KIND_CLUSTES=("${CLUSTER1_NAME}" "${CLUSTER2_NAME}")
for cluster in "${KIND_CLUSTES[@]}"; do
    echo "load image to ${cluster}"
    kind load docker-image quay.io/metallb/controller:${METALLB_VERSION} --name ${cluster}
    kind load docker-image quay.io/metallb/speaker:${METALLB_VERSION} --name ${cluster}
    kind load docker-image istio/proxyv2:${ISTIO_TAG} --name ${cluster}
    kind load docker-image istio/pilot:${ISTIO_TAG} --name ${cluster}

    echo "install metallb in $cluster"
    kubectl create ns metallb-system --kubeconfig="${MEMBER_CLUSTER_KUBECONFIG}" --context="${cluster}"
    util::install_metallb ${MEMBER_CLUSTER_KUBECONFIG} ${cluster}
    kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/${METALLB_VERSION}/manifests/metallb.yaml --kubeconfig="${MEMBER_CLUSTER_KUBECONFIG}" --context="${cluster}"
done

echo "install istio in cluster1"

istioctl install --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER1_NAME}" \
    -f iop/cluster1-eastwest.yaml -y
# config annotation on namespace
kubectl label --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER1_NAME}" namespace istio-system topology.istio.io/network=network1 --overwrite

# expose services
kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER1_NAME}" -f expose-istiod.yaml -nistio-system
kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER1_NAME}" -f expose-services.yaml -nistio-system

CLUSTER1_DISCOVER_ADDRESS=$(kubectl get --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER1_NAME}" svc -nistio-system istio-eastwestgateway  -o jsonpath="{.status.loadBalancer.ingress[0].ip}")

# install helloworld-v1
kubectl label --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER1_NAME}" namespace default istio-injection=enabled --overwrite
kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER1_NAME}" -f https://raw.githubusercontent.com/istio/istio/master/samples/helloworld/helloworld.yaml
kubectl delete deploy helloworld-v2 --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER1_NAME}"

# add cluster2 to cluster1
istioctl x create-remote-secret --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} \
    --context="${CLUSTER2_NAME}" \
    --name=${CLUSTER2_NAME} | \
    kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER1_NAME}" -f - 

echo "install istio in cluster2"

# copy istio-ca-secret from cluster1
kubectl get secret -nistio-system istio-ca-secret -oyaml \
    --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER1_NAME}" | \
    kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER2_NAME}" -f -

istioctl install --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER2_NAME}" \
    -f iop/cluster2-eastwest.yaml -y

CLUSTER2_DISCOVER_ADDRESS=$(kubectl get --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER2_NAME}" svc -nistio-system istio-eastwestgateway  -o jsonpath="{.status.loadBalancer.ingress[0].ip}")

# config annotation on namespace
kubectl label --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER2_NAME}" namespace istio-system topology.istio.io/network=network2 --overwrite

# expose services
kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER1_NAME}" -f expose-istiod.yaml -nistio-system
kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER2_NAME}" -f expose-services.yaml

# add cluster1 to cluster2
istioctl x create-remote-secret --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} \
    --context="${CLUSTER1_NAME}" \
    --name=${CLUSTER1_NAME} | \
    kubectl apply -f - --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER2_NAME}"

# install helloworld-v2
kubectl label --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER2_NAME}" namespace default istio-injection=enabled --overwrite
kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER2_NAME}" -f https://raw.githubusercontent.com/istio/istio/master/samples/helloworld/helloworld.yaml
kubectl delete deploy helloworld-v1 --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${CLUSTER2_NAME}"

function print_success() {
  echo "Multicluster Istio is running."
  echo -e "\nTo manage your clusters, run:"
  echo -e "  export KUBECONFIG=${MEMBER_CLUSTER_KUBECONFIG}"
  echo "Please use 'kubectl config use-context cluster1/cluster2' to switch to the different remote cluster."
}

print_success
