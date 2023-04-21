# Global Ratelimit


## setup envoyproxy/ratelimit

```bash
kubectl apply -n istio-system -f /root/go/src/github.com/zirain/playground/istio/policy/ratelimit/global/rate-limit-service.yaml
```

## apply policy

```
kubectl apply -n istio-system -f /root/go/src/github.com/zirain/playground/istio/policy/ratelimit/global/envoyfilter.yaml
```

## verify

```
# return 429 after 1 req/min
kubeclt exec -it deploy/sleep -- curl httpbin:8000/
# return 429 after 5 req/min
kubeclt exec -it deploy/sleep -- curl httpbin:8000/get
# return 429 after 5 req/min
kubeclt exec -it deploy/sleep -- curl httpbin:8000/status/200
```


## Cleanup

```
kubectl delete -n istio-system -f /root/go/src/github.com/zirain/playground/istio/policy/ratelimit/global/rate-limit-service.yaml
kubectl delete -n istio-system -f /root/go/src/github.com/zirain/playground/istio/policy/ratelimit/global/envoyfilter.yaml
```