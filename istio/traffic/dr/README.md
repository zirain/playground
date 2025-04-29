

For https://github.com/istio/istio/discussions/56079

```
kubectl exec -it exec -it deployments/sleep -- curl "httpbin:8000/drip?code=200&delay=8"
```