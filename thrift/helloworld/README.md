# generater

```
thrift -gen go hello.thrift
```


# run

```
envoy -c envoy.yaml

go run cmd/server/main.go
go run cmd/client/main.go
```