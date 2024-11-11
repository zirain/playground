# IPv6

https://github.com/kubernetes-sigs/kind/issues/1736#issuecomment-660627062

```console
docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPv6Gateway}}{{end}}' envoy-gateway-control-plane
```


```console
socat UDP6-RECVFROM:5300,fork UDP4-SENDTO:127.0.0.53:53 
```