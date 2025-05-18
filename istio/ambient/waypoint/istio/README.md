# Waypoint


```console
istioctl waypoint apply --name simple-http-waypoint
```

```console
httpbin-8495b59fd6-gq47x                1/1     Running   0          78m     10.244.3.16   envoy-gateway-worker2   <none>           <none>
simple-http-waypoint-7b8b74dfb9-q6q97   1/1     Running   0          3m25s   10.244.3.18   envoy-gateway-worker2   <none>           <none>
sleep-5f884b787f-8mrp4                  1/1     Running   0          78m     10.244.1.16   envoy-gateway-worker    <none>           <none>
```

```

# ztunnel-mhhxn
rc.identity="spiffe://cluster.local/ns/default/sa/sleep" dst.addr=10.244.3.18:15008 dst.hbone_addr=10.96.81.28:8000 dst.service="httpbin.default.svc.cluster.local" dst.workload="simple-http-waypoint-7b8b74dfb9-q6q97" dst.namespace="default" dst.identity="spiffe://cluster.local/ns/default/sa/simple-http-waypoint" direction="outbound" bytes_sent=79 bytes_recv=620 duration="17ms"

# simple-http-waypoint-7b8b74dfb9-q6q97
[2025-04-01T07:36:47.565Z] "GET /get HTTP/1.1" 200 - via_upstream - "-" 0 369 7 7 "-" "curl/8.12.1" "cd38105b-41ad-46de-9cae-d440d131d453" "httpbin:8000" "envoy://connect_originate/10.244.3.16:8080" inbound-vip|8000|http|httpbin.default.svc.cluster.local envoy://internal_client_address/ 10.96.81.28:8000 10.244.1.16:37736 - default


# ztunnel-qs5d2
2025-04-01T07:36:52.578960Z     info    access  connection complete     src.addr=10.244.3.18:46796 src.workload="simple-http-waypoint-7b8b74dfb9-q6q97" src.namespace="default" src.identity="spiffe://cluster.local/ns/default/sa/simple-http-waypoint" dst.addr=10.244.3.16:15008 dst.hbone_addr=10.244.3.16:8080 dst.service="httpbin.default.svc.cluster.local" dst.workload="httpbin-8495b59fd6-gq47x" dst.namespace="default" dst.identity="spiffe://cluster.local/ns/default/sa/httpbin" direction="inbound" bytes_sent=565 bytes_recv=156 duration="5008ms"
```


## Install httpbin in another namespace


```console
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: simple-http-waypoint
  namespace: default
spec:
  gatewayClassName: istio-waypoint
  listeners:
  - allowedRoutes:
      namespaces:
        from: All # change this
    name: mesh
    port: 15008
    protocol: HBONE
```


```console
kubectl create ns not-ambient
kubectl apply -f httpbin.yaml -n not-ambient
```

```console
kubectl apply -f networking/not-ambient.yaml
```


```console
kubectl exec -it deploy/sleep -- curl httpbin:8000/get
kubectl exec -it deploy/sleep -- curl httpbin.not-ambient:8000/get
```