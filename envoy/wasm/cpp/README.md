# WASM c++ demo

most of the code come from: https://github.com/proxy-wasm/proxy-wasm-cpp-sdk

## Build sdk

```console
cd proxy-wasm-cpp-sdk
docker build -t wasmsdk:v2 -f Dockerfile-sdk .
```


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
[2022-08-03 09:37:36.508][3814563][warning][wasm] [source/extensions/common/wasm/context.cc:1173] wasm log: [myproject.cc:33]::onLog() onLog id @ b03a4820-9e39-4484-9886-5c29395403a0
[2022-08-03 09:37:36.508][3814563][warning][wasm] [source/extensions/common/wasm/context.cc:1173] wasm log: [myproject.cc:38]::onLog() onLog user @ zirain
[2022-08-03 09:37:36.508][3814563][warning][wasm] [source/extensions/common/wasm/context.cc:1173] wasm log: [myproject.cc:41]::onLog() onLog 3
```