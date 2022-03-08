
envoy -c envoy/ratelimit/basic.yaml --component-log-level filter:debug

fortio load -n 11 127.0.0.1:8000/foo/bar2

curl 127.0.0.1:9901/stats/prometheus | grep http_local_rate_limit