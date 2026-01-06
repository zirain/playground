# Shadow mode in Envoy local rate limiting

```shell
envoy-dev -c envoy/ratelimit/local/shadow-mode/envoy.yaml
```

## Testing shadow mode

```shell
curl -v localhost:8000/svc1
curl -v localhost:8000/svc2
```