
envoy -c envoy/ratelimit/basic.yaml --component-log-level filter:debug,router:debug

envoy -c envoy/ratelimit/multi-route.yaml --component-log-level filter:debug,router:debug

fortio load -n 11 127.0.0.1:8000/svc1/bar2

curl 127.0.0.1:9901/stats/prometheus | grep http_local_rate_limit


# 存在的问题

1. 缺少日志，难以定位
1. `HeaderValueMatch` 缺少 `descriptor_key`生成的 descriptors 不易读
https://github.com/envoyproxy/envoy/issues/20249
1. actions中任意action执行失败后，populateDescriptor返回失败