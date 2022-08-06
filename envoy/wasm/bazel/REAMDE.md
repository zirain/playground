get more info: https://github.com/istio/istio/tree/master/tests/integration/telemetry/stats/prometheus/wasm/wasm_modules/header_injector


## Build wasm

```console
rm -rf myproject.wasm 
docker rm $(docker ps -a | grep wasmsdk | awk '{print $1}')
docker run -v $PWD:/work -w /work  wasmsdk:v2 /build_wasm.sh
```

## Run demo

```console
envoy -c envoy.yaml --component-log-level filter:debug,router:debug;wasm:debug
```

send test traffic:
```console
curl -H "x-real-user: zirain" 127.0.0.1:18000
```

log output:
```console
[2022-08-06 11:34:42.871][419663][warning][wasm] [source/extensions/common/wasm/context.cc:1173] wasm log: [plugin.cc:30]::onLog() request time @ 2022-08-06T03:34:42.870962414+00:00
[2022-08-06 11:34:42.871][419663][warning][wasm] [source/extensions/common/wasm/context.cc:1173] wasm log: [plugin.cc:35]::onLog() request id @ 1d83f88f-3991-48d3-ad47-7fab617b333b
[2022-08-06 11:34:42.871][419663][warning][wasm] [source/extensions/common/wasm/context.cc:1173] wasm log: [plugin.cc:40]::onLog() request user @ zirain
[2022-08-06 11:34:42.871][419663][warning][wasm] [source/extensions/common/wasm/context.cc:1173] wasm log: [plugin.cc:43]::onLog() id 3
```