# Istio build tool

This's a quick way to update istio build tool image:

```
docker build -t ghcr.io/zirain/build-tools:latest .

IMG=zirain/build-tools:latest make lint
```