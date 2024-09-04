```shell
docker run -d -p 5000:5000 --restart=always --name kind-registry registry:2
```

```shell
cat <<EOF | kind create cluster --name istio --config=-
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


get more [here](https://medium.com/@deekonda.ajay/create-your-own-secured-docker-private-registry-with-ssl-6a44539f74b8)


```shell
curl -sS 192.168.3.73:5000/v2/_catalog
curl -sS 192.168.3.73:5000/v2/ratelimit/tags/list

#curl -sS -H 'Accept: application/vnd.docker.distribution.manifest.v2+json' \
curl -sS  -H 'Accept: application/vnd.oci.image.index.v1+json' -o /dev/null \
192.168.3.73:5000/v2/ratelimit/manifests/26f28d78 -w '%header{Docker-Content-Digest}'

curl -sS -X DELETE 192.168.3.73:5000/v2/ratelimit/manifests/sha256:730f0abca15fe886f8786074c0a3b8734e799a5efafc603b9388d43db1684583

rm -rf /var/lib/registry/docker/registry/v2/repositories/ratelimit/
registry garbage-collect -m /etc/docker/registry/config.yml
```