# Envoy Gateway and Istio: test case summary

This document consolidates the checked-in experiments with Envoy Gateway (EG),
Istio sidecars, ambient, waypoints, and multicluster routing. Results describe
these configurations and versions; they are not a compatibility guarantee for
other releases.

| Experiment | Versions | Topology and details |
| --- | --- | --- |
| EG with a sidecar, or with generated Istio certificates | EG v1.10.0, Istio 1.32.0-rc.1 | [Single cluster and primary-remote multicluster](istio-sidecar-multicluster/README.md) |
| EG enrolled in ambient | EG v1.10.0, Istio 1.31.3 | [Ambient multi-primary across two networks](istio-ambient-multicluster/README.md) |
| Ambient and sidecar interoperability | Istio 1.31.3 | [Client/backend matrix with HBONE and 15443 gateways](../istio/multicluster/README.md#ambient-multi-primary-with-sidecar-workloads) |

## Which solution to use

Choose EG's integration to match the backend data plane and the networks it must
reach. The recommendations below follow the tested configurations above.

| Your setup | Recommended solution | Required configuration and limits |
| --- | --- | --- |
| Ambient backends, in one or multiple clusters | Enroll EG proxy pods in ambient | Use `routingType: Service` for remote endpoints and Service waypoints. No Host rewrite or certificate-generation container is needed |
| Sidecar backends, in one or multiple clusters | Inject an Istio sidecar into EG proxy pods | Disable inbound interception, use `routingType: Service`, and rewrite Host to the backend Service FQDN. Across networks, provide a sidecar east-west gateway on 15443 |
| Sidecar backends in one cluster, and EG must run without an Istio sidecar | Use `gen-istio-cert` with EG backend TLS | Configure the client certificate and a `BackendTLSPolicy` per Service. This supplies mTLS but does not provide mesh-wide routing |
| Mixed ambient and sidecar backends within one network | Use the integration matching your intended backend migration | Ambient EG reached both backend types in these tests. An existing sidecar integration can also reach both within the network; retain its Service routing and Host rewrite |
| Mixed ambient and sidecar backends across networks | Standardize the backend data plane, or use separate EG deployments for each backend type | Route ambient backends through ambient EG and sidecar backends through sidecar EG. Both HBONE and 15443 east-west gateways are needed for the two paths; adding gateways does not make cross-network ambient/sidecar interoperability work |

For a new deployment with ambient backends, start with ambient enrollment and
`routingType: Service`. For an existing sidecar mesh, use the sidecar integration
to retain mesh-wide routing. Choose generated certificates when the requirement
is local-cluster mTLS without an EG sidecar, rather than multicluster discovery.

If you need one EG integration to reach every backend across networks, use the
same data plane for EG and those backends. For an ambient migration, move remote
sidecar backends to ambient before sending their traffic through ambient EG.
Separate EG deployments for mixed networks are a recommendation derived from the
individual tested paths; a combined deployment was not tested here.

## EG integration and endpoint discovery

`routingType: Endpoint` makes EG use the local Kubernetes EndpointSlices.
`routingType: Service` makes EG send to the local Service ClusterIP. In the latter
case, an EG sidecar or ztunnel can perform mesh-wide endpoint selection.

| EG integration | `Endpoint` routing | `Service` routing | Cross-cluster behavior |
| --- | --- | --- | --- |
| No sidecar or ambient enrollment, plaintext | Local pods only; STRICT sidecar backends reject plaintext | Local kube-proxy endpoints only; STRICT sidecar backends reject plaintext | No mesh-wide discovery; a remote-only Service fails |
| No sidecar, `gen-istio-cert` and `BackendTLSPolicy` | EG originates mTLS to local endpoints; works with STRICT sidecar backends | Still limited to local kube-proxy endpoints | Certificate provisioning alone does not add remote endpoint discovery or east-west routing |
| Istio sidecar; inbound interception disabled | Tested requests fail with 503: pod-IP traffic falls through to plaintext `PassthroughCluster`, or the local endpoint list is empty | Works with a Service-FQDN Host rewrite; the sidecar selects mesh endpoints | Same-network remote pods are reached directly; remote-network sidecar pods need a 15443 east-west gateway |
| Ambient enrollment | Local endpoints only; bypasses Service waypoints | ztunnel selects mesh-wide endpoints and uses the Service waypoint when configured | Remote ambient backends work through HBONE; remote-network sidecar backends fail in these tests |

For the sidecar integration, rewrite the Host header to
`<service>.<namespace>.svc.cluster.local`. Without it, the tested HTTP requests
miss the sidecar's mesh route and fail against STRICT backends. Ambient routing
uses the destination address and needs no such rewrite.

The certificate-based case uses `gen-istio-cert` v0.6.0, an EG client certificate,
and a `BackendTLSPolicy` for each backend Service. It supplies mTLS credentials,
but EG still needs its own endpoint discovery and routing configuration.

## Cross-network data plane interoperability

The following matrix describes mesh clients in cluster1 reaching HTTP/1.1
backends in cluster2.
Within the same network, all of these data plane combinations worked.

| Client → remote backend | HBONE east-west gateway only | HBONE plus sidecar east-west gateway on 15443 |
| --- | --- | --- |
| Ambient → ambient | 200 | 200 |
| Ambient → ambient with waypoint | 200 | 200 |
| Ambient → sidecar | Connection reset | Connection reset |
| Sidecar → sidecar | 503 `no_healthy_upstream` | 200 |
| Sidecar → ambient, with or without waypoint | 503 `no_healthy_upstream` | 503 `no_healthy_upstream` |

The current ambient multicluster installer installs both gateway types in each
cluster. Observations that sidecar clients cannot reach any remote backend refer
to the earlier HBONE-only configuration.

Sidecars cross networks through mTLS `AUTO_PASSTHROUGH` on 15443. They cannot use
the HBONE-only east-west gateway. Even with 15443 available, istiod excludes the
remote ambient endpoints from the sidecar's endpoint list in this setup.

Ambient clients use double HBONE across networks. The inner CONNECT names the
destination Service. A destination ztunnel handles that authority, but the tested
sidecar's HBONE listener requires a pod-IP destination. Its CONNECT fails, and
the source ztunnel resets the application's connection. Adding 15443 does not
change this path.

## EG results across two ambient clusters

These are the EG-specific outcomes for HTTP/1.1 backends. The in-mesh comparison
client is ambient. Both clusters have matching Service definitions, including
Services whose pods exist only in the other cluster.

| Backend | EG `Endpoint` | EG `Service` | Ambient mesh client |
| --- | --- | --- | --- |
| Local ambient, with or without waypoint | 200; waypoint bypassed | 200 | 200 |
| Remote ambient, with or without waypoint | 503; no local endpoints | 200 | 200 |
| Local sidecar | 200 | 200 | 200 |
| Remote-network sidecar | 503; no local endpoints | 503 from failed tunnel | Reset; curl reports `000` |
| `whereami`, ambient pods in both clusters | Local pods only | Both clusters | Both clusters |
| `whereami`, sidecar pods in both clusters | Local pods only | Local successes mixed with remote failures | Local successes mixed with remote resets |

## Load distribution

| Case | Observation | Interpretation |
| --- | --- | --- |
| Sidecar primary-remote, `whereami` with two pods per cluster | EG `Service`: 26 primary, 27 remote1, 7 remote2 out of 60 requests; remote2 traffic reached one pod | Cross-network `AUTO_PASSTHROUGH` selects a pod per mTLS connection; persistent pools can keep requests on a few pods |
| Same test, mesh client | 30 primary, 19 remote1, 11 remote2; both remote2 pods reached | Different connection picks can produce different distributions |
| Ambient with waypoint, `whereami` | EG `Service`: 17 local, 13 remote out of 30; remote traffic reached one pod. Mesh client: 27 local, 3 remote | Uneven distribution was observed; the experiment did not isolate the cause |

These small samples do not establish expected percentages. In the sidecar test,
the remote gateway also appeared with weight 1 despite representing two pods;
the reason was not investigated. More client connections or shorter connection
lifetimes can create more remote pod selections, but these experiments do not
demonstrate a specific balancing guarantee.

## Policy placement and reproduction

The sidecar multicluster example uses one primary/config cluster and two remotes.
The `h2test` STRICT `PeerAuthentication` is intentionally applied only to the
primary: its istiod distributes the policy to remote sidecars in that namespace.
The ambient example uses independent primaries and applies its namespace policies
in both clusters.

Reproduction commands, manifests, and scripts are in the
[sidecar example](istio-sidecar-multicluster/README.md#run-it) and
[ambient example](istio-ambient-multicluster/README.md#run-it). The scripts print
observations rather than asserting expected results. A printed `000` is not an
HTTP response; check curl's error and proxy logs to distinguish a reset from
other connection failures.
