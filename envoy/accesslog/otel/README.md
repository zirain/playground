# OpenTelemetry ALS

## Setup OTel-collector in VM

``` shell
docker rm -f otel-collector && docker run -d --name otel-collector \
    --network host \
    -v /root/go/src/playground/envoy/access-log/otel-config.yaml/:/etc/otelcol-contrib/config.yaml \
    ghcr.io/open-telemetry/opentelemetry-collector-releases/opentelemetry-collector-contrib:0.51.0
```

## start envoy

```shell
envoy-dev -c envoy/access-log/otel.yaml --component-log-level filter:debug,router:debug
```

## verify 

```shell
curl  127.0.0.1:10000/

curl -H "x-b3-traceid:80f198ee56343ba864fe8b2a57d3eff7" -H "x-b3-spanid:64bdd57d8eda52dd" 127.0.0.1:10000/

docker logs otel-collector
```

## output

```console
2022-06-10T02:49:49.765Z	INFO	loggingexporter/logging_exporter.go:71	LogsExporter	{"#logs": 1}
2022-06-10T02:49:49.765Z	DEBUG	loggingexporter/logging_exporter.go:81	ResourceLog #0
Resource SchemaURL: 
Resource labels:
     -> log_name: STRING(otel_envoy_accesslog)
     -> zone_name: STRING()
     -> cluster_name: STRING()
     -> node_name: STRING()
     -> deploy_type: STRING(binary)
ScopeLogs #0
ScopeLogs SchemaURL: 
InstrumentationScope  
LogRecord #0
ObservedTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2022-06-10 02:49:49.704323941 +0000 UTC
Severity: 
Body: [2022-06-10T02:49:49.704Z] "GET / HTTP/1.1" 200 - direct_response - "-" 0 13 0 - "-" "curl/7.58.0" "03876f8f-01ea-4ea8-a9ff-d00e188097fb" "127.0.0.1:10000" "-" - - 127.0.0.1:10000 127.0.0.1:59962 - -

Trace ID: 
Span ID: 
Flags: 0
```