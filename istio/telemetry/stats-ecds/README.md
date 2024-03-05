# How to use Stats ECDS


## Install Istio

Install Istio with default revision:

```shell
istioctl install -y
```

Install Istio with canary revision:

```shell
install install -y --set revision=canary --set hub=istio --set tag=zirain-dev -f iop.yaml
```

## Setup Application

```shell
kubectl label ns default istio-injection=enabled --overwrite
kubectl apply -f samples/httpbin/httpbin.yaml
kubectl apply -f samples/helloworld/helloworld.yaml
```

```shell
kubectl create ns ns1
kubectl label ns ns1 istio-injection=enabled
kubectl apply -f samples/sleep/sleep.yaml -n ns1
export SLEEP1=$(kubectl get pod -l app=sleep -n ns1 -o jsonpath={.items..metadata.name})
```

```shell
kubectl create ns ns2
kubectl label namespace ns2 istio-injection- istio.io/rev=canary
kubectl apply -f samples/sleep/sleep.yaml -n ns2
export SLEEP2=$(kubectl get pod -l app=sleep -n ns2 -o jsonpath={.items..metadata.name})
```

Check the application is running:

```shell
istioctl x es $SLEEP1 -n ns1 -oprom | grep envoy_cluster_upstream_cx_rx_bytes_total
istioctl x es $SLEEP2 -n ns2 -oprom | grep envoy_cluster_upstream_cx_rx_bytes_total
```

Then create/delete telemetry to trigger full push:

```shell
cat <<EOF | kubectl apply -n istio-system -f -
apiVersion: telemetry.istio.io/v1alpha1
kind: Telemetry
metadata:
  name: mesh-default
  namespace: istio-system
spec:
  accessLogging:
    - providers:
        - name: envoy
EOF
# delete Telemetry API after a while
kubectl delete telemetry -n istio-system mesh-default
```

## Result

| before | after  | delta  |
| ------ | ------ | ------ |
| 180440 | 600845 | 420405 |
| 200063 | 468192 | 268129 |