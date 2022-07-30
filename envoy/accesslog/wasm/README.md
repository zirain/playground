# WASM log


## Build

```console
make build
```

## Run

```
envoy -c envoy.yaml --component-log-level filter:debug,router:debug;wasm:debug
```