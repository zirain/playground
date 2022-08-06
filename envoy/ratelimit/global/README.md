# Ratelimi



```console
envoy-dev -c envoy/ratelimit/global/envoy.yaml --component-log-level filter:debug,router:debug

curl 127.0.0.1:18000/
```

















```console
REGISTRY=zirain VERSION=dev make docker_image
kind load docker-image zirain/ratelimit:dev
```

test stats:
```

kubectl apply -f envoy/ratelimit/global/ef.yaml 

k exec -it sleep-5887ccbb67-gxqxl -- curl -v ratelimit:6070/stats

k exec -it sleep-5887ccbb67-gxqxl -- curl -v ratelimit:6070/rlconfig

k exec -it sleep-5887ccbb67-gxqxl -- curl -v httpbin:8000/get
```