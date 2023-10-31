# Metadata example

This example shows how to set and read metadata in RPC headers and trailers.
Please see
[grpc-metadata.md](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md)
for more information.

## Start the server

```
envoy -c go/grpc/metadata/envoy.yaml -l info
```

```
cd go/grpc/metadata
go run server/main.go
```

## Run the client

```
cd go/grpc/metadata
go run client/main.go --addr 127.0.0.1:18000
```
