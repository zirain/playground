# Envoy Gateway with an Istio sidecar: http2 ports, routingType, multicluster

Envoy Gateway (EG/TEG) runs as edge proxy, with an Istio sidecar injected into its
proxy pods and inbound interception switched off. This is the setup in the Tetrate doc
[Use TEG with Istio](https://docs.tetrate.io/envoy-gateway/howto/eg-and-istio).
These tests check that doc's caveats, on one cluster and on a multicluster mesh.

Tested with Istio 1.32.0-rc.1 and EG v1.10.0, on kind.

## Conclusions

1. **The doc's `http2` caveat blames the wrong hop.** The doc says an `http2` port
   (name prefix or `appProtocol`) "breaks connection between the TEG sidecar and the
   app sidecar". That hop works. The failure is the last hop, inside the app's pod:
   - Istio treats an `http2` port as HTTP/2 and builds the app sidecar's inbound
     cluster with HTTP/2 options, so the sidecar sends cleartext HTTP/2 (h2c) to the app.
   - An app that only speaks HTTP/1.1 rejects it. The app's sidecar logs
     `inbound|8080|| 502 UPE upstream_reset_before_response_started{protocol_error}`.
   - Any mesh client fails the same way, with or without TEG.
   - An app that really speaks h2c (nginx with `http2 on`) works through TEG.

   The rule is Istio's, not TEG's: mark a port HTTP/2 only if the app serves h2c.
   For h2c apps prefer `appProtocol: kubernetes.io/h2c`; Istio treats it as HTTP/2,
   and EG then sends h2c too. EG ignores `http2`.

2. **`routingType: Service` is required.** With the default `Endpoint`, every
   request failed:
   - EG sends to pod IPs. Its sidecar has no mesh route for a pod IP, so it passes
     the connection through in plaintext (`PassthroughCluster`). STRICT mTLS on the
     backend rejects it (`NR filter_chain_not_found`), and EG returns 503.
   - In multicluster, EG reads only its own cluster's EndpointSlices. A Service whose
     pods are all in remote clusters has no endpoints there, so EG answers with a
     `direct_response` 503 and never sends the request.

   With `Service`, EG sends to the ClusterIP, and its sidecar picks endpoints in every
   cluster, like any mesh client.

3. **The Host header must be the Service's FQDN.** Without the doc's `URLRewrite` to
   `<svc>.<ns>.svc.cluster.local`, a host such as `foo.example.com` matches no route in
   the sidecar. The request also goes out in plaintext through `PassthroughCluster`
   and gets 503.

4. **Services that share a pod port share its protocol.** Istio builds one inbound
   cluster per pod port. When an `http2` Service and an `http` Service selected the
   same pod port, the port became HTTP/2 for both, and the `http` Service failed too.

5. **Load to another network is balanced per connection, not per request.**
   - A cluster on another network shows up in the client sidecar as one endpoint: its
     east-west gateway (`172.18.0.200:15443`).
   - That gateway runs in `AUTO_PASSTHROUGH` mode. It routes the mTLS connection by SNI,
     picks one pod per connection, and never sees the requests inside.
   - Envoy keeps one connection pool per worker thread. With `concurrency: 2`, each
     sidecar opened 2 connections to the gateway, so it got at most 2 pod picks.
   - Observed: both of TEG's connections landed on the same remote2 pod, and the mesh
     client's two landed on different pods. That's luck, not a TEG difference.
   - The gateway also has `weight: 1`, like a single pod, though it stands for 2 pods.
     So remote2 got about 1/5 of requests instead of 2/6. I didn't check in istiod why.

   EG keeps long-lived connections to its sidecar, so a few steady connections can stay
   on one remote pod. To spread load, open more connections (higher `concurrency`, more
   EG replicas) or limit connection life with `connectionPool.http.maxRequestsPerConnection`
   in a `DestinationRule`. EG's sidecar is the client here, so it applies the rule.

6. **Without a sidecar, EG only reaches backends in its own cluster.** Multicluster
   works above only because EG's sidecar does the mesh routing: it knows endpoints in
   every cluster and sends to remote2 through its east-west gateway. Without one:
   - **Plain EG** (`nosidecar.yaml`) sends plaintext. STRICT backends reject it (503).
     With PERMISSIVE it works, but only to pods in the primary.
   - **gen-istio-cert** (`gen-istio-cert.yaml`, Option B of the Tetrate doc) gives EG an
     Istio certificate, so EG does mTLS itself, and STRICT backends in the primary accept
     it. The `http2` result is the same as with a sidecar: `http2` port with httpbin 502
     (from the app's sidecar), nginx h2c 200. But EG reads only the primary's
     EndpointSlices, so:
     - a Service whose pods are all remote returns `direct_response` 503;
     - a Service with pods in every cluster (`whereami`) only ever reaches primary pods;
     - `routingType: Service` doesn't help: kube-proxy only knows local pods, and the
       remote-only Services fail with `remote connection failure`.

     Reaching remote2 would also need EG to send mTLS to the east-west gateway with
     Istio's SNI (`outbound_.<port>_._.<svc FQDN>`), which nothing here configures.

   So for multicluster, put the sidecar in (or use ambient, as in
   [istio-mesh-ingress](../istio-mesh-ingress)). gen-istio-cert covers one cluster only.

### Suggested doc wording

Replace the `http2` caveat with:

> Only declare a backend port as HTTP/2 (an `http2` port-name prefix, or
> `appProtocol: http2` / `kubernetes.io/h2c`) if the app serves cleartext HTTP/2.
> Otherwise the app's sidecar sends HTTP/2 to it and requests fail with 502. This is
> standard Istio behavior and isn't specific to TEG. For h2c apps, prefer
> `appProtocol: kubernetes.io/h2c` so TEG also sends HTTP/2.

And add conclusions 2–4 as caveats. The caveat was first written in 2023, ending with
"for reasons I don't understand" (`tetrateio/sites` commit `eaaf3fbe`); the sample app
then was `kong/httpbin`, which also only speaks HTTP/1.1.

## Results

### One cluster

`single-cluster/h2test.yaml`, applied to a cluster with Istio and EG installed.
Backends use STRICT mTLS; each Service has its own Deployment.

| Backend Service port | via TEG (`routingType: Service`) | via mesh client |
| --- | --- | --- |
| `name: http`, httpbin | 200 | 200 |
| `name: http2`, httpbin | 502 | 502 |
| `appProtocol: http2`, httpbin | 502 | 502 |
| `appProtocol: kubernetes.io/h2c`, httpbin | 502 | 502 |
| `name: http2`, nginx speaking h2c | 200, "over HTTP/2.0" | 200 |

With `routingType: Endpoint`, every row is 503.

### Multicluster

Primary-remote from [istio/multicluster](../../istio/multicluster), `ISTIO_NETWORK_MODE=non-flat`:

- **primary:** istiod and EG, on `network1`
- **remote1:** `network1`, reached directly by pod IP
- **remote2:** `network2`, reached through its east-west gateway

Port test (`test.sh`), with each cluster running its own copy of the backends:

| Backend | TEG `Endpoint` | TEG `Service` | mesh client |
| --- | --- | --- | --- |
| `http-*`, httpbin | 503 | 200 in all 3 clusters | 200 |
| `http2-*`, httpbin | 503 | 502 in all 3 clusters | 502 |
| `http2` port, nginx speaking h2c | 503 | 200 over HTTP/2.0, all 3 clusters | 200 |

For remote2, TEG's sidecar sent the request over mTLS to the east-west gateway, and
the 502 came from the app's sidecar in remote2, the same as on one cluster.

Traffic test (`traffic.sh 60`), Service `whereami` with 2 replicas in each cluster:

| Client | primary | remote1 | remote2 |
| --- | --- | --- | --- |
| TEG `Endpoint` | 60 × 503 | | |
| TEG `Service` | 26 | 27 | 7, one pod |
| mesh client | 30 | 19 | 11, both pods |

## Run it

One cluster:

```shell
kubectl apply -f single-cluster/h2test.yaml
kubectl -n envoy-gateway-system port-forward svc/sidecar-service 18091:80 &
curl -s localhost:18091/get -H 'Host: http2.example.com'   # 502
curl -s localhost:18091/    -H 'Host: nginx-http2.example.com'   # 200 over HTTP/2.0
```

The EG proxy Service name depends on EG's deploy mode; find it with
`kubectl -n envoy-gateway-system get svc -l gateway.envoyproxy.io/owning-gateway-name=sidecar-service`.

Multicluster:

```shell
cd ../../istio/multicluster
ISTIO_NETWORK_MODE=non-flat make create-clusters connect-clusters install-istio
K=$PWD/.kube; cd -

helm install eg oci://docker.io/envoyproxy/gateway-helm --version v1.10.0 \
  -n envoy-gateway-system --create-namespace --kubeconfig $K/primary
kubectl --kubeconfig $K/primary wait -n envoy-gateway-system deploy/envoy-gateway \
  --for=condition=Available --timeout=5m

cd multicluster
for c in primary remote1 remote2; do
  kubectl --kubeconfig $K/$c apply -f services.yaml -f workloads-$c.yaml -f whereami-$c.yaml
done
kubectl --kubeconfig $K/primary apply -f primary-only.yaml -f whereami-route.yaml

./test.sh        # port protocol matrix
./traffic.sh 60  # where 60 requests land
```

Both scripts read kubeconfigs from `istio/multicluster/.kube` unless you pass a
directory.

Inspect connections per endpoint in a client sidecar:

```shell
kubectl --kubeconfig $K/primary -n h2test exec meshclient -c istio-proxy -- \
  pilot-agent request GET clusters | grep '^outbound|8000||whereami' | grep -E 'cx_total|rq_total'
```

## Files

| File | Contents |
| --- | --- |
| `single-cluster/h2test.yaml` | Backends on `http`, `http2`, `appProtocol: http2`, `kubernetes.io/h2c` ports, nginx h2c control, two EG gateways with sidecars (`Endpoint`, `Service`), routes. The mesh client was started with `kubectl -n h2test run meshclient --image=curlimages/curl:8.10.1 --command -- sleep 1d` |
| `multicluster/services.yaml` | Services, applied to every cluster |
| `multicluster/workloads-<cluster>.yaml` | httpbin and nginx h2c backends for that cluster |
| `multicluster/primary-only.yaml` | STRICT `PeerAuthentication`, mesh client, EG gateways, routes |
| `multicluster/whereami-*.yaml` | `whereami` Service and pods per cluster, and its route |
| `multicluster/test.sh`, `traffic.sh` | The two tests above |
| `multicluster/nosidecar.yaml` | EG without a sidecar, plaintext to the mesh. Pick the backend with the `x-backend` header |
| `multicluster/gen-istio-cert.yaml` | EG without a sidecar, with gen-istio-cert v0.6.0 and a `BackendTLSPolicy` per Service. Gateway `gen-istio-cert` in namespace `ingress` |
