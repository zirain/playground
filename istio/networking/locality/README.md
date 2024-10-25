
# Locality Load Balancing

1. Create a kind cluster:

```shell
kind create cluster --image=kindest/node:v1.28.0 --config=- <<EOF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
- role: worker
EOF
```

1. Use `topology.kubernetes.io/zone` to label each worker with a zone name:

```shell
kubectl label node kind-worker topology.kubernetes.io/zone=cn-east1
kubectl label node kind-worker2 topology.kubernetes.io/zone=cn-west2
```

```shell
kubectl label node envoy-gateway-worker topology.kubernetes.io/zone=cn-east1 --overwrite
kubectl label node envoy-gateway-worker2 topology.kubernetes.io/zone=cn-west2 --overwrite
```