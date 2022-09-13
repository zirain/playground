# Istio dev scripts

## Setup cluster

```console
cd /root/go/src/istio.io/istio
./samples/kind-lb/setupkind.sh --cluster-name istio --k8s-release 1.24.0

```

## Make all

```console
cd /root/go/src/istio.io/istio
export HUB=istio
export TAG=1.16-dev
KIND_CLUSTER_NAME=istio
make build && make docker.operator && make docker.proxyv2 && make docker.pilot && make docker.install-cni && kind load docker-image istio/proxyv2:$TAG --name ${KIND_CLUSTER_NAME} && kind load docker-image istio/pilot:$TAG --name ${KIND_CLUSTER_NAME} && kind load docker-image istio/operator:$TAG --name ${KIND_CLUSTER_NAME} && kind load docker-image istio/install-cni:$TAG --name ${KIND_CLUSTER_NAME} && unset TAG && unset HUB && docker image prune -f
kubectl delete po --all

```