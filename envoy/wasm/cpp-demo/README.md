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
docker run -v $PWD:/work -w /work  wasmsdk:v2 /build_wasm.sh
```

## Run demo

```console
envoy -c envoy.yaml --component-log-level filter:debug,router:debug;wasm:debug
```

send test traffic:
```console
curl 127.0.0.1:18000
```

log output:
```console
[2022-08-02 18:53:04.280][3617715][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log: onRequestHeaders 3
[2022-08-02 18:53:04.280][3617715][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log: header path /
[2022-08-02 18:53:04.280][3617715][warning][wasm] [source/extensions/common/wasm/context.cc:1173] wasm log: [myproject.cc:25]::onLog() onLog 3
[2022-08-02 18:53:04.280][3617715][info][wasm] [source/extensions/common/wasm/context.cc:1170] wasm log: onDone 3
```