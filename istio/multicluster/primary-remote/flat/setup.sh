#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

# set -x

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

CLUSTER_VERSION=${CLUSTER_VERSION:-"kindest/node:v1.23.4"}
KIND_LOG_FILE=${KIND_LOG_FILE:-"/tmp/istio"}

# create host cluster and member clusters in parallel
# host IP address: script parameter ahead of macOS IP

#prepare for kindClusterConfig
TEMP_PATH=$(mktemp -d)
KIND_CONFIGS_PATH="./kind-configs"
ADDONS_PATH="./addons/"
echo -e "Preparing kindClusterConfig in path: ${TEMP_PATH}"
cp -rf ${KIND_CONFIGS_PATH}/primary.yaml "${TEMP_PATH}"/primary.yaml
cp -rf ${KIND_CONFIGS_PATH}/remote1.yaml "${TEMP_PATH}"/remote1.yaml
cp -rf ${KIND_CONFIGS_PATH}/remote2.yaml "${TEMP_PATH}"/remote2.yaml

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

echo "install istio in primary"
istioctl install --kubeconfig="${MAIN_KUBECONFIG}" -f iop/primary.yaml -y
DISCOVER_ADDRESS=$(kubectl get --kubeconfig=${MAIN_KUBECONFIG} svc -nistio-system istiod -ojsonpath="{ .spec.clusterIP }")

kubectl label --kubeconfig=${MAIN_KUBECONFIG} namespace default istio-injection=enabled
kubectl apply --kubeconfig=${MAIN_KUBECONFIG} -f https://raw.githubusercontent.com/istio/istio/master/samples/httpbin/sample-client/fortio-deploy.yaml
kubectl apply --kubeconfig=${MAIN_KUBECONFIG} -f https://raw.githubusercontent.com/istio/istio/master/samples/sleep/sleep.yaml
# apply helloworld service and dr
kubectl apply --kubeconfig=${MAIN_KUBECONFIG} -f ${ADDONS_PATH}/helloworld/helloworld.yaml

echo "install istio in remote1"
istioctl x create-remote-secret --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} \
    --context="${REMOTE_CLUSTER_1_NAME}" \
    --name=${REMOTE_CLUSTER_1_NAME} | \
    kubectl apply -f - --kubeconfig="${MAIN_KUBECONFIG}"
kubectl get secret -nistio-system istio-ca-secret -oyaml --kubeconfig=${MAIN_KUBECONFIG} | kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_1_NAME}" -f -
istioctl install --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_1_NAME}" \
    --set values.global.remotePilotAddress="${DISCOVER_ADDRESS}" \
    -f iop/remote1.yaml -y

# install helloworld-v1
kubectl label --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_1_NAME}" namespace default istio-injection=enabled
kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_1_NAME}" -f https://raw.githubusercontent.com/istio/istio/master/samples/helloworld/helloworld.yaml
kubectl delete deploy helloworld-v2 --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_1_NAME}"

kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_1_NAME}" -f https://raw.githubusercontent.com/istio/istio/master/samples/addons/prometheus.yaml
export REMOTE1_PROMETHEUS_SVC_IP=$(kubectl get svc prometheus -nistio-system --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_1_NAME}" -o jsonpath='{.spec.clusterIP}')
echo -e "remote1 prometheus endpoint: ${REMOTE1_PROMETHEUS_SVC_IP}"

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

# install helloworld-v1
kubectl label --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_2_NAME}" namespace default istio-injection=enabled
kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_2_NAME}" -f https://raw.githubusercontent.com/istio/istio/master/samples/helloworld/helloworld.yaml
kubectl delete deploy helloworld-v1 --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_2_NAME}"

kubectl apply --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_2_NAME}" -f https://raw.githubusercontent.com/istio/istio/master/samples/addons/prometheus.yaml
export REMOTE2_PROMETHEUS_SVC_IP=$(kubectl get svc prometheus -nistio-system --kubeconfig=${MEMBER_CLUSTER_KUBECONFIG} --context="${REMOTE_CLUSTER_2_NAME}" -o jsonpath='{.spec.clusterIP}')
echo -e "remote2 prometheus endpoint: ${REMOTE2_PROMETHEUS_SVC_IP}"

# install prometheus in primary
kubectl apply --kubeconfig=${MAIN_KUBECONFIG} -f https://raw.githubusercontent.com/istio/istio/master/samples/addons/prometheus.yaml
kubectl apply --kubeconfig=${MAIN_KUBECONFIG} -f https://raw.githubusercontent.com/istio/istio/master/samples/addons/grafana.yaml
cp ${ADDONS_PATH}/prometheus.yaml prometheus-temp.yaml 
sed -i -- "s/REMOTE1_PROMETHEUS_SVC_IP/${REMOTE1_PROMETHEUS_SVC_IP}/g" prometheus-temp.yaml
sed -i -- "s/REMOTE2_PROMETHEUS_SVC_IP/${REMOTE2_PROMETHEUS_SVC_IP}/g" prometheus-temp.yaml
kubectl apply --kubeconfig=${MAIN_KUBECONFIG} -f prometheus-temp.yaml
rm -rf prometheus-temp.yaml

function print_success() {
  echo "Multicluster Istio is running."
  echo -e "\nTo start using your Istio, run:"
  echo -e "  export KUBECONFIG=${MAIN_KUBECONFIG}"
  echo -e "\nTo manage your remote clusters, run:"
  echo -e "  export KUBECONFIG=${MEMBER_CLUSTER_KUBECONFIG}"
  echo "Please use 'kubectl config use-context remote1/remote2' to switch to the different remote cluster."
}

print_success
