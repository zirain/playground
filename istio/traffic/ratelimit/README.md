
```shell
export FORTIO_POD=`kubectl get po | grep fortio | awk '{print $1}'`

# 修改日志级别
istioctl pc log `kubectl get po | grep fortio | awk '{print $1}'` --level filter:debug 

kubectl rollout restart deployment httpbin

kubectl exec -it `kubectl get po | grep fortio | awk '{print $1}'` -- fortio load -n 10  -abort-on 503 httpbin.default.svc.cluster.local:8000/ip

kubectl exec -it `kubectl get po | grep fortio | awk '{print $1}'` -- fortio load -n 10  -abort-on 503 httpbin.default.svc.cluster.local:8000/

kubectl exec -it `kubectl get po | grep fortio | awk '{print $1}'` -- fortio load -n 10  -abort-on 503 httpbin.default.svc.cluster.local:8000/status/200

kubectl exec -it `kubectl get po | grep fortio | awk '{print $1}'` -- fortio load -n 10  -abort-on 503 httpbin:8000/status/200


```