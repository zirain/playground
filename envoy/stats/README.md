# Stats

## start envoy

```shell
envoy -c envoy/stats/envoy.yaml --component-log-level filter:debug,router:debug
```


```shell
curl 127.0.0.1:10000

curl 127.0.0.1:9901/stats/prometheus
```


issue: https://github.com/envoyproxy/envoy/issues/22591