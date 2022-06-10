# Setup OTel-collector in VM

``` shell

docker rm -f otel-collector && docker run -d --name otel-collector \
    --network host \
    -v /root/go/src/playground/envoy/access-log/otel-config.yaml/:/etc/otelcol-contrib/config.yaml \
    ghcr.io/open-telemetry/opentelemetry-collector-releases/opentelemetry-collector-contrib:0.51.0

```

start envoy
```shell
envoy-dev -c envoy/access-log/otel.yaml --component-log-level filter:debug,router:debug
```

verify output
```
curl  127.0.0.1:10000/

docker logs otel-collector
```