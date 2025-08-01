

```shell
kubectl exec -it -n sidecar deploy/sleep -- sh -c 'curl -s https://en.wikipedia.org/wiki/Main_Page | grep -o "<title>.*</title>"'

kubectl exec -it -n sidecar deploy/sleep -- sh -c 'curl -s https://en.wikipedia.org:8443/wiki/Main_Page | grep -o "<title>.*</title>"'
```