
```shell
export FORTIO_POD=`kubectl get po | grep fortio | awk '{print $1}'`


istioctl pc log $HTTPBIN_POD 
export HTTPBIN_POD=`kubectl get po | grep httpbin | awk '{print $1}'`
istioctl pc all $HTTPBIN_POD -oyaml > $HTTPBIN_POD.yaml

kubectl rollout restart deployment httpbin
kubectl exec -it $FORTIO_POD -- fortio load -n 5  -abort-on 503 httpbin.default.svc.cluster.local:8000/ip
kubectl exec -it $FORTIO_POD -- fortio load -n 5  -abort-on 503 httpbin.default.svc.cluster.local:8000/

```