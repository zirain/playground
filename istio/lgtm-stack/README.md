# Setup LGTM stack with Istio (WIP)


## Update addons

```
bash gen.sh
```


## Install istio

```
istioctl install -f iop.yaml -y
```

## Install Otel

```
kubectl apply -f otel.yaml && k rollout restart deploy -nistio-system otel-collector

k logs -l app=opentelemetry -nistio-system -f
```

## Install LGTM stack

```
kubectl apply -f ns.yaml
kubectl apply -f addons/ -n lgtm-stack
```