# local development

https://github.com/howardjohn/local-istio-development


# get istioctl

```console
curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.13.2 TARGET_ARCH=x86_64 sh -
```

# istio devstats

https://istio.teststats.cncf.io/


# PR dashboard

https://k8s-gubernator.appspot.com/pr

# Setup istio e2e cluster

```shell
# MULTI_CLUSTER
prow/integ-suite-kind.sh --topology MULTICLUSTER doc.test.multicluster

# SINGLE_CLUSTER 
prow/integ-suite-kind.sh doc.test.profile_none
```

# How to test istio.io repo using Kind

```
# MACOS
TEST_ENV=kind ADDITIONAL_CONTAINER_OPTIONS="--network host"  make doc.test TEST=tasks/observability/logs/otel-provider
TEST_ENV=kind ADDITIONAL_CONTAINER_OPTIONS="--network host"  make doc.test TEST=tasks/policy-enforcement/rate-limit

export KUBECONFIG=/root/.kube/member1.config:/root/.kube/member2.config:/root/.kube/member3.config
export KUBECONFIG_FILES=(/root/.kube/member1.config /root/.kube/members.config /root/.kube/member3.config)
export KUBE_CONTEXTS=(member1 member2 member3)
TAG="1.13.3" TEST_ENV=kind ADDITIONAL_CONTAINER_OPTIONS="--network host" make doc.test.multicluster TEST=setup/install/external-controlplane
```


# Run istio repo e2e test

```
cd $GOPATH/src/istio.io/istio
go test -tags=integ ./tests/integration/telemetry
```

# Run race test

```bash
go clean -testcache && go test -v -count=1 -race -timeout 30s -run ^TestEDSUnhealthyEndpoints$ istio.io/istio/pilot/pkg/xds
```

# Run benchmark test

```console
go test -benchmem -run=^$ -bench ^BenchmarkListenerGeneration$ istio.io/istio/pilot/pkg/xds
```

# debug ECDS

```shell
kubectl exec -it `kubectl get po -nistio-system| grep istiod | grep Running | awk '{print $1}'` -nistio-system -- pilot-discovery request get '/debug/ecdsz?proxyID=httpbin-74fb669cc6-htg48'
```

```shell
kubectl exec -it `kubectl get po -nistio-system| grep istiod | grep Running | awk '{print $1}'` -nistio-system -- pilot-discovery request get '/debug/config_dump?types=ecds&proxyID=httpbin-74fb669cc6-htg48'
```

# istioctl x metric

```shell
out/linux_amd64/istioctl --log_output_level=debug x metrics httpbin -d 2m
```


```shell
cd $GOPATH/src/istio.io/istio
export HUB=istio
export TAG=1.14-dev
make istioctl && make docker.proxyv2 && make docker.pilot && kind load docker-image istio/proxyv2:$TAG --name istio && kind load docker-image istio/pilot:$TAG --name istio && unset TAG && unset HUB && docker image prune -f


kubectl rollout restart deployment istiod -nistio-system

```


# build

```
cd /root/go/src/istio.io/istio
export HUB=istio
export TAG=1.15-dev
make build && make docker.operator && make docker.proxyv2 && make docker.pilot && make docker.install-cni && kind load docker-image istio/proxyv2:$TAG --name istio && kind load docker-image istio/pilot:$TAG --name istio && kind load docker-image istio/operator:$TAG --name istio && kind load docker-image istio/install-cni:$TAG --name istio && unset TAG && unset HUB && docker image prune -f
kubectl delete po --all
```


# ratelimit

```console
cd ~/go/src/istio.io/ratelimit
make compile
REGISTRY=zirain VERSION=dev make docker_image
kind load docker-image zirain/ratelimit:dev --name istio
kubectl rollout restart deployment ratelimit
```

# grafana

```console
make gen-addons && kubectl apply -f samples/addons/grafana.yaml && kubectl rollout restart deployment grafana -nistio-system

# port-forward
kubectl port-forward -nistio-system pod/`kubectl get po -nistio-system | grep grafana | awk '{print $1}'` 3000:3000 --address 0.0.0.0
```