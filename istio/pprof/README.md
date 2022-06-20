# Analyzing Istio Performance

# Control Plane

port-forward Pilot to local:
```
kubectl port-forward `kubectl get po -nistio-system | grep istiod | awk '{print $1}'` -nistio-system 8080
```


Profile CPU:

```
go tool pprof -http=:8888 localhost:8080/debug/pprof/profile

pprof -web pilot-discovery pprof.pilot-discovery.samples.cpu.001.pb.gz
```