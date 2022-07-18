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
MAIN_KUBECONFIG=${MAIN_KUBECONFIG:-"${KUBECONFIG_PATH}/istio-primary.config"}
MEMBER_CLUSTER_KUBECONFIG=${MEMBER_CLUSTER_KUBECONFIG:-"${KUBECONFIG_PATH}/istio-remotes.config"}
PRIMARY_CLUSTER_NAME=${PRIMARY_CLUSTER_NAME:-"primary"}
REMOTE_CLUSTER_1_NAME=${REMOTE_CLUSTER_1_NAME:-"remote1"}
REMOTE_CLUSTER_2_NAME=${REMOTE_CLUSTER_2_NAME:-"remote2"}
HOST_IPADDRESS=${1:-}

METALLB_VERSION=${METALLB_VERSION:-"v0.10.2"}
CLUSTER_VERSION=${CLUSTER_VERSION:-"kindest/node:v1.23.4"}
KIND_LOG_FILE=${KIND_LOG_FILE:-"/tmp/istio"}

# create host cluster and member clusters in parallel
# host IP address: script parameter ahead of macOS IP

#prepare for kindClusterConfig
TEMP_PATH=$(mktemp -d)
echo -e "Preparing kindClusterConfig in path: ${TEMP_PATH}"
cp -rf primary.yaml "${TEMP_PATH}"/primary.yaml
cp -rf remote1.yaml "${TEMP_PATH}"/remote1.yaml
cp -rf remote2.yaml "${TEMP_PATH}"/remote2.yaml

util::create_cluster "${PRIMARY_CLUSTER_NAME}" "${MAIN_KUBECONFIG}" "${CLUSTER_VERSION}" "${KIND_LOG_FILE}"
util::create_cluster "${REMOTE_CLUSTER_1_NAME}" "${MEMBER_CLUSTER_KUBECONFIG}" "${CLUSTER_VERSION}" "${KIND_LOG_FILE}" "${TEMP_PATH}"/remote1.yaml
util::create_cluster "${REMOTE_CLUSTER_2_NAME}" "${MEMBER_CLUSTER_KUBECONFIG}" "${CLUSTER_VERSION}" "${KIND_LOG_FILE}" "${TEMP_PATH}"/remote2.yaml

util::check_clusters_ready "${MAIN_KUBECONFIG}" "${PRIMARY_CLUSTER_NAME}"
sleep 5s
util::check_clusters_ready "${MEMBER_CLUSTER_KUBECONFIG}" "${REMOTE_CLUSTER_1_NAME}"
sleep 5s
util::check_clusters_ready "${MEMBER_CLUSTER_KUBECONFIG}" "${REMOTE_CLUSTER_2_NAME}"
sleep 5s

# connecting networks between primary, remote1 and remote2 clusters
echo "connect remote1 <-> remote2"
util::connect_kind_clusters "${REMOTE_CLUSTER_1_NAME}" "${MEMBER_CLUSTER_KUBECONFIG}" "${REMOTE_CLUSTER_2_NAME}" "${MEMBER_CLUSTER_KUBECONFIG}" 1

echo "connect primary <-> remote1"
util::connect_kind_clusters "${PRIMARY_CLUSTER_NAME}" "${MAIN_KUBECONFIG}" "${REMOTE_CLUSTER_1_NAME}" "${MEMBER_CLUSTER_KUBECONFIG}" 1

echo "connect primary <-> remote2"
util::connect_kind_clusters "${PRIMARY_CLUSTER_NAME}" "${MAIN_KUBECONFIG}" "${REMOTE_CLUSTER_2_NAME}" "${MEMBER_CLUSTER_KUBECONFIG}" 1

echo "cluster networks connected"

echo "starting install metallb in primary cluster"
kubectl create ns metallb-system --kubeconfig="${MAIN_KUBECONFIG}" --context="${PRIMARY_CLUSTER_NAME}"
util::install_metallb ${MAIN_KUBECONFIG} ${PRIMARY_CLUSTER_NAME}
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/${METALLB_VERSION}/manifests/metallb.yaml --kubeconfig="${MAIN_KUBECONFIG}" --context="${PRIMARY_CLUSTER_NAME}"

echo "starting install metallb in member clusters"
MEMBER_CLUSTERS=(${REMOTE_CLUSTER_1_NAME} ${REMOTE_CLUSTER_2_NAME})
for c in ${MEMBER_CLUSTERS[@]}
do
  echo "install metallb in $c"
  kubectl create ns metallb-system --kubeconfig="${MEMBER_CLUSTER_KUBECONFIG}" --context="${c}"
  util::install_metallb ${MEMBER_CLUSTER_KUBECONFIG} ${c}
  kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/${METALLB_VERSION}/manifests/metallb.yaml --kubeconfig="${MEMBER_CLUSTER_KUBECONFIG}" --context="${c}"
done

echo "install istio in primary"
istioctl install --kubeconfig="${MAIN_KUBECONFIG}" -f iop/primary.yaml -y
# config annotation on namespace
kubectl label --kubeconfig=${MAIN_KUBECONFIG} namespace istio-system topology.istio.io/network=network1 --overwrite
# install eastwest gateway
istioctl install --kubeconfig="${MAIN_KUBECONFIG}" -f iop/primary-eastwest.yaml -y 
# expose istiod
kubectl apply --kubeconfig=${MAIN_KUBECONFIG} -f expose-istiod.yaml
# expose services
kubectl apply --kubeconfig=${MAIN_KUBECONFIG} -f expose-services.yaml

kubectl apply --kubeconfig=${MAIN_KUBECONFIG} -f helloworld/helloworld.yaml

DISCOVER_ADDRESS=$(kubectl get --kubeconfig=${MAIN_KUBECONFIG} svc -nistio-system istio-eastwestgateway  -o jsonpath="{.status.loadBalancer.ingress[0].ip}")

kubectl label --kubeconfig=${MAIN_KUBECONFIG} namespace default istio-injection=enabled --overwrite
kubectl apply --kubeconfig=${MAIN_KUBECONFIG} -f https://raw.githubusercontent.com/istio/istio/master/samples/httpbin/sample-client/fortio-deploy.yaml
kubectl apply --kubeconfig=${MAIN_KUBECONFIG} -f https://raw.githubusercontent.com/istio/istio/master/samples/sleep/sleep.yaml

echo "install istio in remote1"
istioctl x create-remote-secret --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} \
    --context="${REMOTE_CLUSTER_1_NAME}" \
    --name=${REMOTE_CLUSTER_1_NAME} | \
    kubectl apply -f - --kubeconfig="${MAIN_KUBECONFIG}"
# create istio ca secret
kubectl get secret -nistio-system istio-ca-secret -oyaml --kubeconfig=${MAIN_KUBECONFIG} | kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_1_NAME}" -f -

istioctl install --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_1_NAME}" \
    --set values.global.remotePilotAddress="${DISCOVER_ADDRESS}" \
    -f iop/remote1.yaml -y
# config annotation on namespace
kubectl label --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_1_NAME}" namespace istio-system topology.istio.io/network=network1 --overwrite
# install eastwest gateway
istioctl install --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_1_NAME}" \
    --set values.global.remotePilotAddress="${DISCOVER_ADDRESS}" \
    -f iop/remote1-eastwest.yaml -y

# expose services
kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_1_NAME}" -f expose-services.yaml

# install helloworld-v1
kubectl label --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_1_NAME}" namespace default istio-injection=enabled --overwrite
kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_1_NAME}" -f https://raw.githubusercontent.com/istio/istio/master/samples/helloworld/helloworld.yaml
kubectl delete deploy helloworld-v2 --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_1_NAME}"

echo "install istio in remote2"
istioctl x create-remote-secret --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} \
    --context="${REMOTE_CLUSTER_2_NAME}" \
    --name=${REMOTE_CLUSTER_2_NAME} | \
    kubectl apply -f - --kubeconfig="${MAIN_KUBECONFIG}"
kubectl get secret -nistio-system istio-ca-secret -oyaml \
    --kubeconfig=${MAIN_KUBECONFIG} | \
    kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_2_NAME}" -f -
istioctl install --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_2_NAME}" \
    --set values.global.remotePilotAddress="${DISCOVER_ADDRESS}" \
    -f iop/remote2.yaml -y
# config annotation on namespace
kubectl label --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_2_NAME}" namespace istio-system topology.istio.io/network=network2 --overwrite
# install eastwest gateway
istioctl install --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_2_NAME}" \
    --set values.global.remotePilotAddress="${DISCOVER_ADDRESS}" \
    -f iop/remote2-eastwest.yaml -y

# expose services
kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_2_NAME}" -f expose-services.yaml

# install helloworld-v2
kubectl label --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_2_NAME}" namespace default istio-injection=enabled --overwrite
kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_2_NAME}" -f https://raw.githubusercontent.com/istio/istio/master/samples/helloworld/helloworld.yaml
kubectl delete deploy helloworld-v1 --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_2_NAME}"

function print_success() {
  echo "Multicluster Istio is running."
  echo -e "\nTo start using your Istio, run:"
  echo -e "  export KUBECONFIG=${MAIN_KUBECONFIG}"
  echo -e "\nTo manage your remote clusters, run:"
  echo -e "  export KUBECONFIG=${MEMBER_CLUSTER_KUBECONFIG}"
  echo "Please use 'kubectl config use-context remote1/remote2' to switch to the different remote cluster."
}

print_success
