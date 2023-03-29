

```
envoy-dev -c envoy/accesslog/cel-command/envoy.yaml --component-log-level filter:debug,router:debug
```


curl 127.0.0.1:10000/cluster1
