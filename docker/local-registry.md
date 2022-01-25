```shell
docker run -d -p 5000:5000 --restart=always --name kind-registry registry:2
```

```shell
cat <<EOF | kind create cluster --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:5000"]
    endpoint = ["http://kind-registry:5000"]
EOF
```

```shell
docker network connect "kind" "kind-registry"
```