# Ratelimit

```console
REGISTRY=zirain VERSION=dev make docker_image
kind load docker-image zirain/ratelimit:dev
```

test stats:
```
k exec -it sleep-5887ccbb67-97lgm -- curl ratelimit:6070/stats

k exec -it sleep-5887ccbb67-97lgm -- curl ratelimit:6070/rlconfig

k exec -it sleep-5887ccbb67-97lgm -- curl httpbin:8000/get
```