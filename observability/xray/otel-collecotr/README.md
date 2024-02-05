# EKS with X-Ray

This will not working, see more details [here](https://github.com/istio/istio/issues/36599#issuecomment-1744600462).

## Create Cluster

```shell
eksctl create cluster \
    --region=us-east-2 \
    --name appmeshtest-zirain \
    --nodes-min 2 \
    --nodes-max 3 \
    --nodes 2 \
    --auto-kubeconfig \
    --full-ecr-access \
    --appmesh-access
```

```shell
export KUBECONFIG=~/.kube/eksctl/clusters/appmeshtest-zirain
```


## Enable ADOT addon

### Install cert-manager

```shell
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.2/cert-manager.yaml
```

### Install ADOT

```shell
aws eks describe-addon-configuration --addon-name adot --addon-version v0.90.0-eksbuild.1
```

```shell
aws eks create-addon --addon-name adot --addon-version v0.90.0-eksbuild.1 --cluster-name appmeshtest-zirain
```

Verify ADOT status:

```shell
aws eks describe-addon --addon-name adot --cluster-name appmeshtest-zirain
```

Install Collector:

```shell
kubectl create ns observability
kubectl apply -f otel-collector.yaml
```

## Install Sample APP

```yaml
kubectl label namespace default istio-injection=enabled --overwrite
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.20/samples/sleep/sleep.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.20/samples/httpbin/httpbin.yaml
```

## Output

```json
{
    "Id": "1-71f86e63-7c9ad592296688596eea6ee8",
    "Duration": 0.003,
    "LimitExceeded": false,
    "Segments": [
        {
            "Id": "cf77a1e7d50cfad8",
            "Document": {
                "id": "cf77a1e7d50cfad8",
                "name": "egress httpbin:8000",
                "start_time": 1707149433.975948,
                "trace_id": "1-71f86e63-7c9ad592296688596eea6ee8",
                "end_time": 1707149433.9784808,
                "fault": false,
                "error": false,
                "throttle": false,
                "http": {
                    "request": {
                        "url": "http://httpbin:8000/get",
                        "method": "GET"
                    },
                    "response": {
                        "status": 0,
                        "content_length": 0
                    }
                },
                "aws": {
                    "xray": {
                        "auto_instrumentation": false
                    }
                },
                "metadata": {
                    "default": {
                        "istio.namespace": "default",
                        "downstream_cluster": "-",
                        "http.protocol": "HTTP/1.1",
                        "otel.resource.service.name": "sleep.default",
                        "response_size": "533",
                        "response_flags": "-",
                        "component": "proxy",
                        "zone": "us-east-2c",
                        "istio.canonical_service": "sleep",
                        "upstream_cluster": "outbound|8000||httpbin.default.svc.cluster.local",
                        "istio.canonical_revision": "latest",
                        "peer.address": "192.168.82.49",
                        "istio.mesh_id": "cluster.local",
                        "request_size": "0",
                        "guid:x-request-id": "36f189be-0e4c-9187-acb5-b7dfdb552328",
                        "user_agent": "curl/8.6.0",
                        "node_id": "sidecar~192.168.82.49~sleep-58f5c6b977-tbpz6.default~default.svc.cluster.local",
                        "upstream_cluster.name": "outbound|8000||httpbin.default.svc.cluster.local"
                    }
                }
            }
        },
        {
            "Id": "17adc8f2065d73ce",
            "Document": {
                "id": "17adc8f2065d73ce",
                "name": "httpbin.default",
                "start_time": 1707149433.9763231,
                "trace_id": "1-71f86e63-7c9ad592296688596eea6ee8",
                "end_time": 1707149433.978102,
                "parent_id": "cf77a1e7d50cfad8",
                "fault": false,
                "error": false,
                "throttle": false,
                "http": {
                    "request": {
                        "url": "http://httpbin:8000/get",
                        "method": "GET"
                    },
                    "response": {
                        "status": 0,
                        "content_length": 0
                    }
                },
                "aws": {
                    "xray": {
                        "auto_instrumentation": false
                    }
                },
                "metadata": {
                    "default": {
                        "istio.namespace": "default",
                        "downstream_cluster": "-",
                        "http.protocol": "HTTP/1.1",
                        "otel.resource.service.name": "httpbin.default",
                        "response_size": "533",
                        "response_flags": "-",
                        "component": "proxy",
                        "zone": "us-east-2c",
                        "istio.canonical_service": "httpbin",
                        "upstream_cluster": "inbound|80||",
                        "istio.canonical_revision": "v1",
                        "peer.address": "192.168.82.49",
                        "istio.mesh_id": "cluster.local",
                        "request_size": "0",
                        "guid:x-request-id": "36f189be-0e4c-9187-acb5-b7dfdb552328",
                        "user_agent": "curl/8.6.0",
                        "node_id": "sidecar~192.168.65.79~httpbin-86869bccff-5dd9m.default~default.svc.cluster.local",
                        "upstream_cluster.name": "inbound|80||"
                    }
                }
            }
        }
    ]
}
```


## Cleanup

```shell
kubectl delete otelcol --all -n observability
kubectl delete ns observability
aws eks delete-addon --addon-name adot --cluster-name appmeshtest-zirain
kubectl delete -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.2/cert-manager.yaml
eksctl delete cluster --region=us-east-2 --name appmeshtest-zirain
```


## Troubleshooting

```console
2024-02-05T15:50:27.429Z	debug	awsxrayexporter@v0.90.1/awsxray.go:69	request: {
  TraceSegmentDocuments: ["{\"name\":\"httpbin.default\",\"id\":\"05cb84e99ab001d4\",\"start_time\":1707148223.302142,\"trace_id\":\"1-737e899b-f5f8a759cb22b2d90e7f3478\",\"end_time\":1707148223.305716,\"http\":{\"request\":{\"method\":\"GET\",\"url\":\"http://httpbin:8000/get\"},\"response\":{\"status\":0,\"content_length\":0}},\"fault\":false,\"error\":false,\"throttle\":false,\"aws\":{\"xray\":{\"auto_instrumentation\":false}},\"metadata\":{\"default\":{\"component\":\"proxy\",\"downstream_cluster\":\"-\",\"guid:x-request-id\":\"1fabc68b-17fc-9f50-9ee1-1eb59a6fff11\",\"http.protocol\":\"HTTP/1.1\",\"istio.canonical_revision\":\"v1\",\"istio.canonical_service\":\"httpbin\",\"istio.mesh_id\":\"cluster.local\",\"istio.namespace\":\"default\",\"node_id\":\"sidecar~192.168.65.79~httpbin-86869bccff-5dd9m.default~default.svc.cluster.local\",\"otel.resource.service.name\":\"httpbin.default\",\"peer.address\":\"192.168.82.49\",\"request_size\":\"0\",\"response_flags\":\"-\",\"response_size\":\"533\",\"upstream_cluster\":\"inbound|80||\",\"upstream_cluster.name\":\"inbound|80||\",\"user_agent\":\"curl/8.6.0\",\"zone\":\"us-east-2c\"}},\"parent_id\":\"4f827cc3b3c9c945\"}\n"]
}	{"kind": "exporter", "data_type": "traces", "name": "awsxray"}
2024-02-05T15:50:27.448Z	debug	awsxrayexporter@v0.90.1/awsxray.go:72	response error	{"kind": "exporter", "data_type": "traces", "name": "awsxray", "error": "AccessDeniedException: User: arn:aws:sts::354280132276:assumed-role/eksctl-appmeshtest-zirain-nodegrou-NodeInstanceRole-UZyQ6kMAIaou/i-0490401a7d73f5c20 is not authorized to perform: xray:PutTraceSegments because no identity-based policy allows the xray:PutTraceSegments action\n\tstatus code: 403, request id: 11fc5319-c0cd-491b-a942-3e593e0c4972"}
```


Need attach policy:

```shell
eksctl utils associate-iam-oidc-provider --region=us-east-2 --cluster=appmeshtest-zirain --approve

eksctl create iamserviceaccount --name x-ray-collector --namespace observability \
    --cluster appmeshtest-zirain \
    --attach-policy-arn arn:aws:iam::aws:policy/AWSXRayDaemonWriteAccess \
    --override-existing-serviceaccounts \
    --approve
```