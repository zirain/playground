# Building multi-platform images

```shell
docker buildx create --name mybuilder --bootstrap --use

docker buildx build --platform linux/amd64,linux/arm64 -t zirain/curl .
```