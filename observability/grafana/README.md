# Grafana Stack

## Grafana

```shell
helm upgrade -i --create-namespace grafana grafana/grafana --namespace monitoring -f observability/grafana/grafana.yaml
```


## Alloy eBPF

https://grafana.com/docs/pyroscope/latest/configure-client/grafana-alloy/ebpf/setup-kubernetes/


```shell
helm upgrade -i alloy-ebpf grafana/alloy -f observability/grafana/ebpf.yaml -n monitoring --create-namespace
```

## Pyroscope

```shell
helm upgrade -i pyroscope grafana/pyroscope -f observability/grafana/pyroscope.yaml -n monitoring --create-namespace
```
