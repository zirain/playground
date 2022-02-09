
# Display Metadata
an simple wasm plugin uesd to verify [WasmPlugin.VmConfig](https://github.com/istio/api/pull/2227)



# Login ghcr

```shell
echo $GHCR_PAT | docker login ghcr.io -u zirain --password-stdin
```

# Build

```shell
make docker.push
```