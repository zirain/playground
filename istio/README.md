# local development

https://github.com/howardjohn/local-istio-development

# istio devstats

https://istio.teststats.cncf.io/


# PR dashboard

https://k8s-gubernator.appspot.com/pr


# How to test istio.io repo using Kind

```
TEST_ENV=kind ADDITIONAL_CONTAINER_OPTIONS="--network host"  make doc.test TEST=tasks/observability/logs/otel-provider
```


# Run istio repo e2e test

```
cd $GOPATH/src/istio.io/istio
go test -tags=integ ./tests/integration/telemetry
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