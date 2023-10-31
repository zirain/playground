# gostats

test on:
```console
go version go1.18.2 linux/amd64
```

```console
go install golang.org/x/perf/cmd/benchstat@latest
```


```
git co master
go test -benchmem -run=^$ -bench ^BenchmarkListenerGeneration$ istio.io/istio/pilot/pkg/xds | tee -a old.txt

git co accesslog-cache
go test -benchmem -run=^$ -bench ^BenchmarkListenerGeneration$ istio.io/istio/pilot/pkg/xds | tee -a new.txt


benchstat old.txt new.txt
```