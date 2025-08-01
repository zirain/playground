

```shell
kubectl exec -it -n sidecar deploy/sleep -- sh -c 'curl -s https://en.wikipedia.org/wiki/Main_Page | grep -o "<title>.*</title>"'

kubectl exec -it -n sidecar deploy/sleep -- sh -c 'curl -s https://en.wikipedia.org:8443/wiki/Main_Page | grep -o "<title>.*</title>"'
```


```shell
kubectl exec -it -n sidecar deploy/sleep -- sh -c 'curl -s https://map.baidu.com | grep -o "<title>.*</title>"'
kubectl exec -it -n sidecar deploy/sleep -- sh -c 'curl -s https://map.baidu.com:8443 | grep -o "<title>.*</title>"'
```