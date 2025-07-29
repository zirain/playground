# gRPC gateway demo

```shell
fortio grpcping 172.18.0.201:80

fortio load -a -grpc -ping -c 2 -s 4 172.18.0.201:80
```