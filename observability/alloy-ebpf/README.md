# Alloy eBPF (WIP)

https://grafana.com/docs/pyroscope/latest/configure-client/grafana-alloy/ebpf/setup-kubernetes/


```shell
helm upgrade -i alloy-ebpf grafana/alloy -f observability/alloy-ebpf/ebpf.yaml -n monitoring --create-namespace
```