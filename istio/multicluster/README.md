# Istio Multicluster

## Primary remote 

### Same network

```shell
make create-clusters connect-clusters install-istio install-app
```


### Different networks

```shell
ISTIO_NETWORK_MODE=non-flat make create-clusters connect-clusters install-istio install-app
```


## Ambient multi-primary

### Different networks

Two primary clusters (`cluster1` on `network1`, `cluster2` on `network2`) running ambient mode,
connected through one east-west gateway per cluster, serving HBONE (15008) for ambient and
15443 for sidecar workloads. A shared root CA is generated into `.certs`.

> [!IMPORTANT]
> The single east-west gateway only works on **Istio 1.32+**. On older versions you need
> two gateways per cluster; see [East-west gateway](#east-west-gateway).

```shell
ISTIO_MC_MODE=ambient-multi-primary ISTIO_NETWORK_MODE=non-flat make up
```

Verify cross-cluster traffic (should return both `v1` and `v2`):

```shell
for i in $(seq 10); do
  kubectl --kubeconfig .kube/cluster1 exec -n ambient deploy/sleep -- curl -s helloworld:5000/hello
done
```

Services are shared across clusters (`serviceScopeConfigs` in the IOPs) when they are:
- in a namespace labeled `istio.io/dataplane-mode=ambient`, unless the service is labeled `istio.io/global=false`, or
- labeled `istio.io/global=true`, in any namespace.

#### East-west gateway

> [!IMPORTANT]
> **Istio 1.32+ only.** One gateway serving both data planes needs
> [istio/istio#62442](https://github.com/istio/istio/pull/62442) and
> `AMBIENT_ENABLE_MULTI_NETWORK=true` (set in the IOPs). On older versions, use two
> gateways, as described below.

**Istio 1.32+: one gateway.**
[eastwest/cluster1.yaml](ambient-multi-primary/non-flat/eastwest/cluster1.yaml) (and
`cluster2.yaml`) defines one `istio-eastwestgateway`, class `istio-east-west`, with two
listeners:

| Listener | Port | Protocol | Used by |
| --- | --- | --- | --- |
| `mesh` | 15008 | HBONE | ztunnel (ambient), double HBONE |
| `cross-network` | 15443 | TLS `Passthrough` → `AUTO_PASSTHROUGH` | sidecars, SNI-routed mTLS |

**Before 1.32: two gateways.** Keep the HBONE `istio-eastwestgateway` with only the
`mesh` listener, and add a second Gateway with class `istio` and only the
`cross-network` listener, for example `istio-eastwestgateway-sidecar`. Each gets its own
Deployment and LoadBalancer. Why:

- **Sidecars can't use an HBONE-only gateway.** istiod skips HBONE-only gateways when it
  builds sidecar endpoints ("Sidecar proxies cannot talk to ambient E/W gateway for now",
  `selectNetworkGateways` in `pilot/pkg/xds/endpoints/ep_filters.go`) and drops the
  remote endpoints. Without a 15443 listener, sidecar clients get
  `503 no_healthy_upstream` for everything in the other network.
- **The single-gateway config is accepted but broken (tested on 1.31.3).** Both
  listeners are `Programmed`, the Service exposes 15443, and istiod points sidecars at
  it. But the `istio-east-west` gateway runs as `proxy waypoint`, and istiod only gives
  it waypoint clusters (`inbound-vip|<port>|tcp|<svc>`). The 15443 filter chains forward
  to `outbound_.<port>_._.<svc>` clusters, which don't exist there, so every connection
  is dropped with `NC` (no cluster) and sidecar → sidecar across networks returns 503.

ztunnel only uses the HBONE listener. istiod finds gateways by the
`topology.istio.io/network` label, not by name. Neither setup makes ambient → sidecar or
sidecar → ambient work across networks; see
[Ambient multi-primary with sidecar workloads](#ambient-multi-primary-with-sidecar-workloads).

## Cleanup

```shell
make down
```

## Tips

### How to install Istio multicluster manually

- Install istio in primary cluster
- Expose istiod service
- Create remote secret with `istioctl x create-remote-secret --server https://172.18.0.3:6443 --name remote > remote.kubeconfig`
  `172.18.0.3` is ip of the docker container
- Install remote istio with discovery address `kubectl get --kubeconfig=${MAIN_KUBECONFIG} svc -nistio-system istio-eastwestgateway  -o jsonpath="{.status.loadBalancer.ingress[0].ip}"`
## Ambient multi-primary with sidecar workloads

Results for a mesh with both ambient and sidecar workloads, across networks (Istio
1.31.3, `ambient-multi-primary/non-flat`). Backends: `amb` (ambient), `amb-wp` (ambient
+ waypoint), `sc` (sidecar), deployed in both clusters; the manifests are in
[envoy-gateway/istio-ambient-multicluster](../../envoy-gateway/istio-ambient-multicluster).
Clients run in cluster1.

| Client (cluster1) → backend (cluster2) | HBONE E/W gateway only | HBONE + 15443 E/W gateways |
| --- | --- | --- |
| ambient → ambient / ambient + waypoint | 200 | 200 |
| ambient → sidecar | fails (reset) | fails (reset) |
| sidecar → sidecar | 503 `no_healthy_upstream` | 200 |
| sidecar → ambient / ambient + waypoint | 503 `no_healthy_upstream` | 503 `no_healthy_upstream` |

Inside one network, every combination works.

### Why requests fail

The two failures look different because different proxies detect them:

- **503** comes from the client's sidecar, an L7 proxy, which has no route to cluster2
  and answers with its own HTTP error.
- **reset** comes from ztunnel, an L4 proxy. When its tunnel fails it can't send an HTTP
  response, so it closes the connection: curl exits with code 56 and prints `000`.
  An L7 client in front of ztunnel turns this into its own 503; Envoy Gateway, for
  example, returns `503 ... reset reason: connection termination`.

#### 503: sidecar client → ambient pod in cluster2

The request never leaves the client pod.

1. **istiod gives the client's sidecar no cluster2 endpoints.** For an endpoint on
   another network, istiod only routes a sidecar through a gateway if the endpoint
   looks like a sidecar workload: the `tlsMode: istio` label that sidecar injection adds
   (`isMtlsEnabled` in `pilot/pkg/xds/endpoints/ep_filters.go`). Ambient pods don't have
   it, so istiod skips them without logging anything.
2. **The sidecar has nowhere to send the request.** The endpoint list is empty:

   ```shell
   istioctl --kubeconfig .kube/cluster1 proxy-config endpoints sidecarclient.sc \
     --cluster 'outbound|8000||http-c2.amb.svc.cluster.local'
   ```

3. **The client's own Envoy answers** `503 UH no_healthy_upstream`, with upstream host `-`.

Even if istiod didn't skip these pods, it wouldn't work. The 15443 gateway would pass
the sidecar's Istio mTLS straight to the pod's app port, and on an ambient pod ztunnel
only accepts mTLS as HBONE on port 15008.

With only the HBONE gateway, sidecar clients get the same 503 for **every** cluster2
backend, sidecars included. Sidecars cross networks with mTLS to an `AUTO_PASSTHROUGH`
gateway on 15443, and istiod drops HBONE-only gateways for sidecars ("Sidecar proxies
cannot talk to ambient E/W gateway for now", `selectNetworkGateways` in the same file),
then the remote endpoints. istiod logs
`Workload ... on network network2 has no reachable E/W gateway for this proxy, skipping`.

#### Reset: ambient client → sidecar pod in cluster2

The request reaches cluster2 and fails at the destination sidecar.

1. **The app's TCP connection succeeds locally.** istio-cni redirects curl's connection
   to the node's ztunnel, which accepts it.
2. **ztunnel tunnels to the Service name.** Across networks it uses double HBONE: an
   outer tunnel to cluster2's HBONE east-west gateway, and an inner, end-to-end mTLS
   tunnel to the destination. ztunnel in cluster1 doesn't know cluster2's pod IPs, so
   the inner request is `CONNECT http-c2.sc.svc.cluster.local:8000`.
3. **The gateway forwards it unchanged.** It can't decrypt the inner tunnel. It picks a
   pod and passes the tunnel to that pod's port 15008, here the sidecar.
4. **The sidecar can't use a hostname.** Its 15008 listener (`connect_terminate`) copies
   the `:authority` into the original destination (`connect_authority` filter,
   `pilot/pkg/xds/filters/filters.go`) and hands the stream to the internal listener
   `main_internal`. That listener only matches `<pod IP>/32` on the target port, so
   nothing matches and the sidecar answers the `CONNECT` with a 503. Its debug log
   (`istioctl proxy-config log <pod> --level conn_handler:debug,connection:debug`):

   ```
   ':method', 'CONNECT'
   ':authority', 'http-c2.sc.svc.cluster.local:8000'
   'x-forwarded-network', 'network1'
   closing connection from 10.20.2.5:57262: no matching filter chain found
   immediate connect error: Invalid argument|remote address:envoy://main_internal/
   ```

   `10.20.2.5` is cluster2's east-west gateway. The access log shows
   `"CONNECT - HTTP/2" 503 UF ... immediate_connect_error:_Invalid_argument`.
5. **ztunnel resets the connection.** ztunnel in cluster1 gets the 503 for its tunnel
   (`error="http status: 503 ..."` in its log), can't turn it into an HTTP response,
   and closes the connection it had accepted. curl sees a reset.

The 15443 gateway doesn't help here: ztunnel only uses HBONE gateways.

#### Why the other cases work

- **Ambient → ambient in cluster2:** the destination node's ztunnel accepts a
  Service-name `CONNECT` and picks the pod itself.
- **Sidecar → sidecar in cluster2 (with the 15443 gateway):** the sidecar sends mTLS to
  the 15443 gateway, naming the Service in the SNI (`outbound_.8000_._.<svc FQDN>`). The
  gateway picks a pod and passes the connection to the sidecar's ordinary inbound port,
  so nothing needs to resolve a Service name. The scripts serve 15443 from the same
  gateway as HBONE (Istio 1.32+ only); see [East-west gateway](#east-west-gateway).

Istio master still skips sidecars for double HBONE ("sidecars don't support double-HBONE
yet"), so the same limits are expected in newer versions.

**In short:** in a multi-network ambient mesh, a workload in another network is
reachable only from a client with the same data plane. Sidecar clients need a 15443
listener, which the scripts add to every cluster's east-west gateway. To mix freely across networks, move workloads to ambient.

Reproduce, with the backends applied:

```shell
kubectl --kubeconfig .kube/cluster1 -n sc run sidecarclient --image=curlimages/curl:8.10.1 --command -- sleep 1d
kubectl --kubeconfig .kube/cluster1 -n amb run meshclient --image=curlimages/curl:8.10.1 --command -- sleep 1d
for t in http-c2.amb http-c2.amb-wp http-c2.sc; do
  kubectl --kubeconfig .kube/cluster1 -n amb exec meshclient -- curl -s -m 5 -o /dev/null -w "ambient  $t %{http_code}\n" http://$t:8000/get
  kubectl --kubeconfig .kube/cluster1 -n sc exec sidecarclient -c sidecarclient -- curl -s -m 5 -o /dev/null -w "sidecar  $t %{http_code}\n" http://$t:8000/get
done
```
