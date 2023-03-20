# http

```shell
export HTTPBIN_POD=`kubectl get po | grep httpbin | awk '{print $1}'`
export FORTIO_POD=`kubectl get po | grep fortio | awk '{print $1}'`
export SLEEP_POD=`kubectl get po | grep sleep | awk '{print $1}'`

kubectl exec -it deploy/sleep -c sleep -- curl httpbin:8000/get
```


```shell
kubectl exec -it $HTTPBIN_POD -cistio-proxy -- curl 127.0.0.1:15020/stats/prometheus | grep istio_requests_total
```

## envoyfilter
```console
istio_requests_total{response_code="200",reporter="destination",source_workload="sleep",source_workload_namespace="default",source_principal="spiffe://cluster.local/ns/default/sa/sleep",source_app="sleep",source_version="unknown",source_cluster="Kubernetes",destination_workload="httpbin",destination_workload_namespace="default",destination_principal="spiffe://cluster.local/ns/default/sa/httpbin",destination_app="httpbin",destination_version="v1",destination_service="httpbin.default.svc.cluster.local",destination_service_name="httpbin",destination_service_namespace="default",destination_cluster="Kubernetes",request_protocol="http",response_flags="-",grpc_response_status="",connection_security_policy="mutual_tls",source_canonical_service="sleep",destination_canonical_service="httpbin",source_canonical_revision="latest",destination_canonical_revision="v1"} 1
```
## telemetry
```console
istio_requests_total{response_code="200",reporter="destination",source_workload="sleep",source_workload_namespace="default",source_principal="spiffe://cluster.local/ns/default/sa/sleep",source_app="sleep",source_version="unknown",source_cluster="Kubernetes",destination_workload="httpbin",destination_workload_namespace="default",destination_principal="spiffe://cluster.local/ns/default/sa/httpbin",destination_app="httpbin",destination_version="v1",destination_service="httpbin.default.svc.cluster.local",destination_service_name="httpbin",destination_service_namespace="default",destination_cluster="Kubernetes",request_protocol="http",response_flags="-",grpc_response_status="",connection_security_policy="mutual_tls",source_canonical_service="sleep",destination_canonical_service="httpbin",source_canonical_revision="latest",destination_canonical_revision="v1"} 1
```


# tcp

```shell
kubectl exec -it $HTTPBIN_POD -cistio-proxy -- curl 127.0.0.1:15020/stats/prometheus | grep istio_tcp_connections_opened_total
```

## envoyfilter
```console
istio_tcp_connections_opened_total{reporter="destination",source_workload="sleep",source_workload_namespace="default",source_principal="spiffe://cluster.local/ns/default/sa/sleep",source_app="sleep",source_version="unknown",source_cluster="Kubernetes",destination_workload="httpbin",destination_workload_namespace="default",destination_principal="spiffe://cluster.local/ns/default/sa/httpbin",destination_app="httpbin",destination_version="v1",destination_service="httpbin.default.svc.cluster.local",destination_service_name="httpbin",destination_service_namespace="default",destination_cluster="Kubernetes",request_protocol="tcp",response_flags="-",connection_security_policy="mutual_tls",source_canonical_service="sleep",destination_canonical_service="httpbin",source_canonical_revision="latest",destination_canonical_revision="v1"} 1
```
## telemetry
```console
istio_tcp_connections_opened_total{reporter="destination",source_workload="sleep",source_workload_namespace="default",source_principal="spiffe://cluster.local/ns/default/sa/sleep",source_app="sleep",source_version="unknown",source_cluster="Kubernetes",destination_workload="httpbin",destination_workload_namespace="default",destination_principal="spiffe://cluster.local/ns/default/sa/httpbin",destination_app="httpbin",destination_version="v1",destination_service="httpbin.default.svc.cluster.local",destination_service_name="httpbin",destination_service_namespace="default",destination_cluster="Kubernetes",request_protocol="tcp",response_flags="-",connection_security_policy="mutual_tls",source_canonical_service="sleep",destination_canonical_service="httpbin",source_canonical_revision="latest",destination_canonical_revision="v1"} 1
```
