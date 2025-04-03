# Sandwitch Demo


```console
k get po -owide
NAME                                          READY   STATUS    RESTARTS   AGE   IP            NODE                    NOMINATED NODE   READINESS GATES
httpbin-8495b59fd6-gq47x                      1/1     Running   0          18m   10.244.3.16   envoy-gateway-worker2   <none>           <none>
simple-http-waypoint-istio-6f76c4df85-wsg24   1/1     Running   0          18m   10.244.3.17   envoy-gateway-worker2   <none>           <none>
sleep-5f884b787f-8mrp4                        1/1     Running   0          18m   10.244.1.16   envoy-gateway-worker    <none>           <none>
```
```console
k get svc -owide
NAME                         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                        AGE   SELECTOR
httpbin                      ClusterIP   10.96.81.28     <none>        8000/TCP                       38m   app=httpbin
kubernetes                   ClusterIP   10.96.0.1       <none>        443/TCP                        41m   <none>
simple-http-waypoint-istio   ClusterIP   10.96.83.29     <none>        15021/TCP,8000/TCP,15008/TCP   28m   gateway.networking.k8s.io/gateway-name=simple-http-waypoint
sleep                        ClusterIP   10.96.136.163   <none>        80/TCP                         38m   app=sleep
```

```console
k get po -n istio-system -owide
NAME                      READY   STATUS    RESTARTS   AGE   IP            NODE                          NOMINATED NODE   READINESS GATES
istio-cni-node-7t2d6      1/1     Running   0          12m   172.18.0.2    envoy-gateway-worker          <none>           <none>
istio-cni-node-p2q28      1/1     Running   0          12m   172.18.0.4    envoy-gateway-control-plane   <none>           <none>
istio-cni-node-rzkpq      1/1     Running   0          12m   172.18.0.3    envoy-gateway-worker2         <none>           <none>
istiod-599cc89f76-wzgvj   1/1     Running   0          12m   10.244.1.12   envoy-gateway-worker          <none>           <none>
ztunnel-9t6cf             1/1     Running   0          12m   10.244.0.6    envoy-gateway-control-plane   <none>           <none>
ztunnel-mhhxn             1/1     Running   0          12m   10.244.1.11   envoy-gateway-worker          <none>           <none>
ztunnel-qs5d2             1/1     Running   0          12m   10.244.3.6    envoy-gateway-worker2         <none>           <none>
```

```console
# ztunnel-mhhxn@envoy-gateway-worker
2025-04-01T06:22:31.259071Z     info    access  connection complete     
src.addr=10.244.1.16:55374 
src.workload="sleep-5f884b787f-8mrp4" 
src.namespace="default" 
src.identity="spiffe://cluster.local/ns/default/sa/sleep" 
dst.addr=10.244.3.17:15008 
dst.hbone_addr=10.96.81.28:8000 
dst.service="httpbin.default.svc.cluster.local" 
dst.workload="simple-http-waypoint-istio-6f76c4df85-wsg24" 
dst.namespace="default" 
dst.identity="spiffe://cluster.local/ns/default/sa/simple-http-waypoint-istio" 
direction="outbound" 
bytes_sent=79 bytes_recv=1514 duration="9ms"

# ztunnel-qs5d2@envoy-gateway-worker2
2025-04-01T06:22:31.258928Z     info    access  connection complete     
src.addr=10.244.1.16:44758 
src.workload="sleep-5f884b787f-8mrp4" 
src.namespace="default" 
src.identity="spiffe://cluster.local/ns/default/sa/sleep" 
dst.addr=10.244.3.17:15008 
dst.hbone_addr=10.96.81.28:8000 
dst.service="simple-http-waypoint-istio.default.svc.cluster.local" 
dst.workload="simple-http-waypoint-istio-6f76c4df85-wsg24" 
dst.namespace="default" 
dst.identity="spiffe://cluster.local/ns/default/sa/simple-http-waypoint-istio" 
direction="inbound" 
bytes_sent=1514 bytes_recv=79 duration="7ms"

# simple-http-waypoint-istio-6f76c4df85-wsg24@envoy-gateway-worker2
[2025-04-01T06:22:31.251Z] "GET /get HTTP/1.1" 200 - via_upstream - "-" 0 1221 3 3 "10.244.1.16" "curl/8.12.1" "27e324bb-9f5d-4353-90db-f759723fd57b" "httpbin:8000" "10.244.3.16:8080" outbound|8000||httpbin.default.svc.cluster.local 10.244.3.17:38844 10.244.3.17:8000 10.244.1.16:49357 - default.httpbin-httproute.0

# ztunnel-qs5d2@envoy-gateway-worker2
2025-04-01T06:22:36.258584Z     info    access  connection complete     
src.addr=10.244.3.17:38844 
src.workload="simple-http-waypoint-istio-6f76c4df85-wsg24" 
src.namespace="default" 
src.identity="spiffe://cluster.local/ns/default/sa/simple-http-waypoint-istio" 
dst.addr=10.244.3.16:15008 
dst.hbone_addr=10.244.3.16:8080 
dst.workload="httpbin-8495b59fd6-gq47x" 
dst.namespace="default" 
dst.identity="spiffe://cluster.local/ns/default/sa/httpbin" 
direction="outbound"
bytes_sent=904 bytes_recv=1418 duration="5006ms"

# ztunnel-qs5d2@envoy-gateway-worker2
2025-04-01T06:22:36.258760Z     info    access  connection complete     
src.addr=10.244.3.17:60152 
src.workload="simple-http-waypoint-istio-6f76c4df85-wsg24" 
src.namespace="default" 
src.identity="spiffe://cluster.local/ns/default/sa/simple-http-waypoint-istio" 
dst.addr=10.244.3.16:15008 
dst.hbone_addr=10.244.3.16:8080 
dst.service="httpbin.default.svc.cluster.local" 
dst.workload="httpbin-8495b59fd6-gq47x" 
dst.namespace="default" 
dst.identity="spiffe://cluster.local/ns/default/sa/httpbin" 
direction="inbound" 
bytes_sent=1418 bytes_recv=904 duration="5004ms"
```