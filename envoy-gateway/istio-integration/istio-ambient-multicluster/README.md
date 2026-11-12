# Envoy Gateway in an ambient multicluster mesh

Envoy Gateway (EG) on cluster1 of an Istio ambient multi-primary mesh. The EG proxy
pods are enrolled in ambient with a pod label, as in [istio-mesh-ingress](../istio-mesh-ingress).
The backends run in both clusters, in three namespaces:

| Namespace | Data plane |
| --- | --- |
| `amb` | ambient, no waypoint |
| `amb-wp` | ambient, with a waypoint for its Services |
| `sc` | sidecar |

Compare with the sidecar-based setup in [istio-sidecar-multicluster](../istio-sidecar-multicluster).

Tested with Istio 1.31.3 and EG v1.10.0 on kind. Two clusters from
[istio/multicluster](../../istio/multicluster) in `ambient-multi-primary/non-flat` mode:
cluster1 on `network1`, cluster2 on `network2`, joined by HBONE east-west gateways.
The installer also adds a sidecar east-west gateway on port 15443 in each cluster.

## Conclusions

1. **`routingType: Service` is required to reach cluster2.** With `Endpoint`, EG only
   sees cluster1's EndpointSlices:
   - a Service with pods only in cluster2 returns 503;
   - `whereami`, with pods in both clusters, only ever reaches cluster1 pods.

   With `Service`, EG sends to the ClusterIP. ztunnel sees the Service, uses the
   mesh-wide endpoints and tunnels to cluster2 through its east-west gateway.
   `Endpoint` also skips the waypoint, because ztunnel only uses the waypoint for
   traffic sent to a Service.

2. **No sidecar, no gen-istio-cert, no Host rewrite.** ztunnel does mTLS for EG, and
   routes by destination address, not by Host header. The routes use plain
   `*.example.com` hosts.

3. **`http2` ports only matter where an L7 proxy sits next to the app.**
   - `amb`, no waypoint: ztunnel is L4 and ignores the port protocol. The `http2` port
     with httpbin returns 200, and nginx sees HTTP/1.1, as EG sent it.
   - `amb-wp`, waypoint: the waypoint treats the port as HTTP/2 and sends h2c to the
     app. httpbin returns 502, nginx h2c returns 200 over HTTP/2.0. The same rule as
     sidecars. With `Endpoint`, the waypoint is skipped, so it returns 200.
   - `sc` in cluster1: the app's sidecar sends h2c. httpbin 502, nginx 200 over HTTP/2.0.

   The rule from the sidecar tests still holds: mark a port HTTP/2 only if the app serves
   h2c. It's not EG-specific.

4. **Sidecar backends in another network are unreachable from ambient clients.** Every
   client failed to reach `sc` pods in cluster2: EG in both modes and the ambient mesh
   client. `whereami.sc` failed whenever ztunnel picked a cluster2 pod.
   - Across networks ztunnel uses double HBONE: an outer tunnel to cluster2's east-west
     gateway, and an inner, end-to-end mTLS tunnel to the destination. ztunnel in
     cluster1 doesn't know cluster2's pod IPs, so the inner `CONNECT` names the Service:
     `:authority: http-c2.sc.svc.cluster.local:8000`. The gateway can't rewrite it,
     because the inner tunnel is encrypted end to end; it just forwards it to the pod's
     port 15008.
   - A sidecar's HBONE listener (`connect_terminate`) expects `<pod IP>:<target port>`
     there. Its `connect_authority` filter copies `%REQ(:AUTHORITY)%` into the original
     destination (`pilot/pkg/xds/filters/filters.go`), then hands the stream to the
     internal listener `main_internal`, whose filter chains match
     `10.20.2.14/32` port `8080`.
   - A hostname isn't an address, so the destination is never set. Debug log of the
     sidecar: `closing connection ...: no matching filter chain found`, then
     `immediate connect error: Invalid argument|remote address:envoy://main_internal/`,
     and a 503 back through the tunnel.
   - On the same network, ztunnel sends `CONNECT <pod IP>:8080` straight to the sidecar,
     so it works. A ztunnel at the destination resolves a Service-name authority
     itself, so ambient pods in cluster2 work.

   So it's an Istio limitation (sidecars can't be double-HBONE destinations), not an EG
   one. The code is the same on Istio master. Sidecar backends in EG's own network work;
   for other networks, enroll the backend in ambient.

   **With only HBONE east-west gateways, sidecar clients can't reach cluster2.**
   In that configuration, a sidecar client in cluster1
   (`sidecarclient` in `sc`) reached every cluster1 backend, but got
   `503 UH no_healthy_upstream` for every cluster2 backend, ambient or sidecar.
   `whereami` only reached cluster1 pods. Its sidecar has no cluster2 endpoints:
   - The ambient east-west gateway only listens on HBONE 15008, not the mTLS port 15443
     that sidecars cross networks with.
   - istiod drops HBONE-only gateways for sidecars ("Sidecar proxies cannot talk to
     ambient E/W gateway for now", `selectNetworkGateways` in
     `pilot/pkg/xds/endpoints/ep_filters.go`), then drops the remote endpoints. Its log:
     `Workload whereami on network network2 has no reachable E/W gateway for this proxy, skipping`.

   The current installer adds a separate mTLS `AUTO_PASSTHROUGH` east-west gateway
   on port 15443 in each cluster. With both gateways, sidecar clients can reach
   sidecar backends in cluster2. Sidecar → ambient and ambient → sidecar traffic
   across networks still fails; the 15443 gateway doesn't resolve those limits.
   See the [cross-network client matrix](../../istio/multicluster/README.md#ambient-multi-primary-with-sidecar-workloads).
   The [sidecar EG setup](../istio-sidecar-multicluster) can therefore reach sidecar
   backends across networks through the 15443 gateway. Enroll EG in ambient to reach
   ambient backends across networks, as in this example.

5. **Load to cluster2 isn't spread evenly.** For `whereami` in `amb-wp`, the 30 requests
   via EG `Service` reached one cluster2 pod; the mesh client reached cluster2 only 3
   times. Traffic to another network goes through the east-west gateway, which picks a
   pod per tunnel, not per request. I didn't dig further.

## Results

`./test.sh 30`. For backends, 200/502/503 is the HTTP status; `000` means the
connection was reset. For nginx, the text says which cluster answered and the HTTP
version the app received.

### `amb`: ambient, no waypoint

| Backend | EG `Endpoint` | EG `Service` | mesh client |
| --- | --- | --- | --- |
| `http-c1` | 200 | 200 | 200 |
| `http2-c1`, httpbin | 200 | 200 | 200 |
| `nginx-h2c-c1` | 200, HTTP/1.1 | 200, HTTP/1.1 | 200 |
| `http-c2` | 503 | 200 | 200 |
| `http2-c2`, httpbin | 503 | 200 | 200 |
| `nginx-h2c-c2` | 503 | 200 c2, HTTP/1.1 | 200 |
| `whereami` ×30 | 30 to c1 | 11 c1, 19 c2 | 14 c1, 16 c2 |

### `amb-wp`: ambient with a waypoint

| Backend | EG `Endpoint` | EG `Service` | mesh client |
| --- | --- | --- | --- |
| `http-c1` | 200 | 200 | 200 |
| `http2-c1`, httpbin | 200 | 502 | 502 |
| `nginx-h2c-c1` | 200, HTTP/1.1 | 200, HTTP/2.0 | 200 |
| `http-c2` | 503 | 200 | 200 |
| `http2-c2`, httpbin | 503 | 502 | 502 |
| `nginx-h2c-c2` | 503 | 200 c2, HTTP/2.0 | 200 |
| `whereami` ×30 | 30 to c1 | 17 c1, 13 c2 (one pod) | 27 c1, 3 c2 |

### `sc`: sidecar

| Backend | EG `Endpoint` | EG `Service` | mesh client |
| --- | --- | --- | --- |
| `http-c1` | 200 | 200 | 200 |
| `http2-c1`, httpbin | 502 | 502 | 502 |
| `nginx-h2c-c1` | 200, HTTP/2.0 | 200, HTTP/2.0 | 200 |
| `http-c2` | 503 | 503 | 000 |
| `http2-c2` | 503 | 503 | 000 |
| `nginx-h2c-c2` | 503 | 503 | 000 |
| `whereami` ×30 | 30 to c1 | 21 c1, 9 × 503 | 15 c1, 15 failed |

## Run it

```shell
cd ../../istio/multicluster
ISTIO_MC_MODE=ambient-multi-primary ISTIO_NETWORK_MODE=non-flat \
  make create-clusters connect-clusters install-istio
K=$PWD/.kube; cd -

# install-istio already applied the Gateway API CRDs with kubectl, so helm's CRD
# install conflicts. Apply EG's CRDs server-side and install the chart without them.
helm template eg-crds oci://docker.io/envoyproxy/gateway-crds-helm --version v1.10.0 \
  --set crds.gatewayAPI.enabled=true --set crds.gatewayAPI.channel=experimental \
  --set crds.envoyGateway.enabled=true |
  kubectl --kubeconfig $K/cluster1 apply --server-side --force-conflicts -f -
helm install eg oci://docker.io/envoyproxy/gateway-helm --version v1.10.0 \
  -n envoy-gateway-system --create-namespace --skip-crds --kubeconfig $K/cluster1
kubectl --kubeconfig $K/cluster1 wait -n envoy-gateway-system deploy/envoy-gateway \
  --for=condition=Available --timeout=5m

kubectl --kubeconfig $K/cluster1 apply -f backends-cluster1.yaml
kubectl --kubeconfig $K/cluster2 apply -f backends-cluster2.yaml
kubectl --kubeconfig $K/cluster1 apply -f gateway-cluster1.yaml

./test.sh 30
```

`test.sh` reads kubeconfigs from `istio/multicluster/.kube` unless you pass a
directory as the second argument.

Check that the EG proxies are in ambient and the controller is not:

```shell
istioctl --kubeconfig $K/cluster1 ztunnel-config workloads | grep envoy-gateway-system
```

## Files

| File | Contents |
| --- | --- |
| `backends-cluster1.yaml`, `backends-cluster2.yaml` | Namespaces `amb`, `amb-wp`, `sc` with STRICT mTLS, Services (same in both clusters), that cluster's httpbin, nginx h2c and `whereami` pods, and the `amb-wp` waypoint. Services in `sc` are labeled `istio.io/global=true`, because only ambient namespaces are global by default |
| `gateway-cluster1.yaml` | Two EG gateways with the proxy enrolled in ambient (`routingType` `Endpoint` and `Service`), PERMISSIVE `PeerAuthentication` for the proxies, an HTTPRoute per backend (host `<svc>.<ns>.example.com`), and the `meshclient` pod in `amb` |
| `test.sh` | The tests above |
