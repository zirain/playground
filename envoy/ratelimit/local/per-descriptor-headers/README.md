# Shadow mode in Envoy local rate limiting

```shell
envoy-dev -c envoy/ratelimit/local/per-descriptor-headers/envoy.yaml -l debug --component-log-level filter:debug,config:debug
```

## Testing shadow mode

```shell
curl -v localhost:8000/svc1
curl -v localhost:8000/svc2
```

```shell
curl -v localhost:8001/svc1
curl -v localhost:8001/svc2
```