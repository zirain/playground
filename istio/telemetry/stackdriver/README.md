# Stackdriver with Istio

## Install (fake) Stackdriver adapter

```shell
kubectl apply -f istio/telemetry/stackdriver/gce-metadata-server.yaml -n istio-system
kubectl apply -f istio/telemetry/stackdriver/fake-sd-server-delpoyment.yaml -n istio-system
```

### Install Istio

```shell
istioctl install -f istio/telemetry/stackdriver/iop.yaml -y
```