# Backend Utilization Load Balancing (ORCA)

`DestinationRule` can use backend-reported utilization to weight endpoints, added in [istio/api#3750](https://github.com/istio/api/pull/3750).
Endpoints that report higher utilization receive fewer requests. It maps to Envoy's
[client_side_weighted_round_robin](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/load_balancers#weighted-round-robin) policy.

> The control-plane support landed in [istio/istio#61289](https://github.com/istio/istio/pull/61289) and ships in Istio 1.32.
> Everything below was verified on `1.32.0-rc.1`.

## API

`backendUtilization` sits next to `simple` and `consistentHash` under `loadBalancer`; only one of them can be set.

```yaml
apiVersion: networking.istio.io/v1
kind: DestinationRule
metadata:
  name: orca
spec:
  host: orca.default.svc.cluster.local
  trafficPolicy:
    loadBalancer:
      backendUtilization:
        weightStabilizationPeriod: 5s
        weightExpirationPeriod: 3m
        weightUpdatePeriod: 1s
        errorUtilizationPenaltyPercent: 100
        metricNamesForComputingUtilization:
          - named_metrics.foo
```

| Field                                | Envoy field                               | Default                   | Meaning                                                                                                                                     |
| ------------------------------------ | ----------------------------------------- | ------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| `weightStabilizationPeriod`          | `blackout_period`                         | 10s                       | After an endpoint's first valid report, wait this long before trusting its weight. Restarts if the endpoint stops reporting and comes back. |
| `weightExpirationPeriod`             | `weight_expiration_period`                | 3m                        | If an endpoint sends no valid report for this long, its weight is dropped and it goes back to the default weight.                           |
| `weightUpdatePeriod`                 | `weight_update_period`                    | 1s                        | How often weights are recomputed. Minimum 100ms.                                                                                            |
| `errorUtilizationPenaltyPercent`     | `error_utilization_penalty` (value / 100) | 0 (off)                   | Extra penalty for endpoints that return errors. `100` weights the error rate 1.0x.                                                          |
| `metricNamesForComputingUtilization` | `metric_names_for_computing_utilization`  | `application_utilization` | ORCA metrics used as utilization; the max of them is taken. Map entries use `<field>.<key>`, e.g. `named_metrics.foo`.                      |

## How it works

The split is decided by the **client** sidecar. Backends only report how loaded they are.

```
 sleep pod                                             orca-busy / orca-idle pods
┌───────────────────────────────┐                     ┌──────────────────────────┐
│ curl ──► sidecar (outbound)   │ ── request ──────►  │ sidecar ──► nginx        │
│          cluster: orca        │                     │                          │
│          LB = CSWRR           │ ◄── response ─────  │ endpoint-load-metrics:   │
│   ① parse ORCA header         │                     │  TEXT application_       │
│   ② store per-endpoint stats  │                     │  utilization=0.9,        │
│   ③ every weightUpdatePeriod: │                     │  rps_fractional=100      │
│      recompute weights        │                     └──────────────────────────┘
│   ④ pick next endpoint by     │
│      weighted round robin     │
└───────────────────────────────┘
```

### 1. Backends send an ORCA load report

ORCA (Open Request Cost Aggregation) is a standard format for a backend to report its own load.
This policy uses the per-request form: a response header that Envoy parses.

| Format | Example                                                                            |
| ------ | ---------------------------------------------------------------------------------- |
| Text   | `endpoint-load-metrics: TEXT application_utilization=0.9,rps_fractional=100`       |
| JSON   | `endpoint-load-metrics: JSON {"application_utilization":0.9,"rps_fractional":100}` |
| Binary | `endpoint-load-metrics-bin: <base64 of OrcaLoadReport proto>`                      |

Real apps fill these from live measurements (e.g. gRPC's ORCA server interceptor or custom middleware).
The demo uses nginx with fixed values so the result is predictable.

### 2. The client sidecar computes a weight per endpoint

```
weight = qps / (utilization + (errors_per_sec / qps) * errorUtilizationPenalty)
```

- `qps` comes from `rps_fractional`.
- `utilization` is the max of `metricNamesForComputingUtilization` if set, otherwise `application_utilization`, falling back to `cpu_utilization`.
- A report with `qps` or `utilization` of 0 or missing is ignored.

Endpoints without a usable weight (new, still stabilizing, or expired) get a default weight, which the API documents as equal weight.
That is why traffic starts out roughly 50/50.

Envoy then runs weighted round robin over the weights, so the split is smooth rather than random.

### 3. Generated Envoy config

The `DestinationRule` becomes this cluster `load_balancing_policy`:

```yaml
lb_policy: LOAD_BALANCING_POLICY_CONFIG
load_balancing_policy:
  policies:
  - typed_extension_config:
      name: envoy.load_balancing_policies.client_side_weighted_round_robin
      typed_config:
        "@type": type.googleapis.com/envoy.extensions.load_balancing_policies.client_side_weighted_round_robin.v3.ClientSideWeightedRoundRobin
        blackout_period: 5s
        weight_expiration_period: 180s
        weight_update_period: 1s
        error_utilization_penalty: 0  # always set; Envoy's own default is 1.0
```

## Demo

[cswrr.yaml](cswrr.yaml) deploys one `orca` Service backed by two nginx Deployments, each returning a fixed ORCA header:

| Endpoint  | qps | utilization | weight | expected share |
| --------- | --- | ----------- | ------ | -------------- |
| orca-idle | 100 | 0.1         | 1000   | ≈ 90%          |
| orca-busy | 100 | 0.9         | ~111   | ≈ 10%          |

1. Enable sidecar injection and deploy the backends and a client:

```shell
kubectl label ns default istio-injection=enabled
kubectl apply -f cswrr.yaml
kubectl apply -f ../../apps/http/sleep.yaml
```

1. Check the client's cluster uses the new policy:

```shell
istioctl pc cluster deploy/sleep --fqdn orca.default.svc.cluster.local -o json | grep -A10 client_side_weighted_round_robin
```

If nothing is printed, your istiod is older than 1.32.

1. Send traffic and count which pod answered:

```shell
kubectl exec deploy/sleep -- sh -c 'for i in $(seq 1 200); do curl -s orca:8080; done' | sort | uniq -c
```

A 200-request loop takes about 1 second, so loops run within `weightStabilizationPeriod` (5s) of the client's first request are split exactly 50/50.
Wait a few seconds and run it again. Observed results:

```
# within the first 5s           # after 5s
 100 orca-busy-...               19 orca-busy-...
 100 orca-idle-...              181 orca-idle-...
```

1. Check the computed weights. `istioctl pc ep` shows the runtime CSWRR weight, not the EDS weight from istiod, so read it while traffic is flowing:

```shell
istioctl pc ep deploy/sleep --cluster "outbound|8080||orca.default.svc.cluster.local" -o yaml | grep -E "address: [0-9]|weight"
```

`orca-idle` shows `1000` (100 / 0.1) and `orca-busy` shows `111`.
If every endpoint shows `weight: 1`, no endpoint has a valid weight. Either no reports have arrived yet, or the last one is older than `weightExpirationPeriod`.

1. Swap the `UTILIZATION` values of the two Deployments. After the pods restart and the weights stabilize, the split reverses (observed 180 `orca-busy` / 20 `orca-idle`):

```shell
kubectl set env deploy/orca-busy UTILIZATION=0.1
kubectl set env deploy/orca-idle UTILIZATION=0.9
```

## Locality failover

`failover` and `failoverPriority` work with `backendUtilization`; `distribute` and zone-aware LB do not.

Istio implements failover by assigning a priority to each endpoint group in the per-proxy load assignment, not through the cluster LB config.
Envoy's CSWRR worker is built on `RoundRobinLoadBalancer`, so it keeps normal priority handling:

1. Envoy picks a priority level based on how healthy each level is.
2. Within that level, CSWRR spreads requests by the ORCA weights.

Envoy rejects `load_balancing_policy` combined with the locality fields of `common_lb_config`, so istiod drops `locality_weighted_lb_config` and zone-aware config for this policy.
`kubectl apply` prints a warning for any `localityLbSetting`, even though failover still works:

```
Warning: backendUtilization is not compatible with localityLbSetting or zoneAwareLbSetting; zone-aware routing and locality weighting will not be applied
```

| Setting                                              | Works with `backendUtilization`? |
| ---------------------------------------------------- | -------------------------------- |
| `localityLbSetting.failover`                         | ✅ uses priorities                |
| `localityLbSetting.failoverPriority`                 | ✅ uses priorities                |
| `trafficDistribution: PreferClose` (forced failover) | ✅ uses priorities                |
| `localityLbSetting.distribute`                       | ❌ locality weights are dropped   |
| Zone-aware LB                                        | ❌ dropped                        |

```yaml
apiVersion: networking.istio.io/v1
kind: DestinationRule
metadata:
  name: orca
spec:
  host: orca.default.svc.cluster.local
  trafficPolicy:
    loadBalancer:
      backendUtilization:
        weightStabilizationPeriod: 5s
      localityLbSetting:
        enabled: true
        failoverPriority:
          - topology.kubernetes.io/zone
    outlierDetection:
      consecutive5xxErrors: 1
      interval: 1s
      baseEjectionTime: 1m
      maxEjectionPercent: 100
```

### Verifying failover

The kind cluster has no zone labels, so the test used the pod label `version` instead:

1. Set `failoverPriority: [version]` and label the client pod `version: busy`.
   The client's `/clusters` output then shows `orca-busy` at `priority::0` and `orca-idle` at `priority::1`, and all 200 requests go to `orca-busy`.
1. Make `orca-busy` return 503.
   After the first 503 it is marked `/failed_outlier_check`, and the other 199 requests go to `orca-idle`.

```shell
kubectl exec deploy/sleep -c istio-proxy -- pilot-agent request GET clusters | grep orca | grep -E "priority|health_flags"
```

Things to watch:

- **Outlier detection is required.** Without `outlierDetection`, istiod doesn't enable failover and traffic never moves to a lower priority.
- **Set `maxEjectionPercent`.** Envoy only ejects a host if `(ejected + 1) / hosts ≤ maxEjectionPercent`, and the default is 10%. With fewer than 10 hosts nothing is ever ejected, so failover never happens. This applies to every LB policy, not just this one.
- **ORCA doesn't trigger failover.** Only outlier ejection or unhealthy endpoints move traffic to the next priority. A local endpoint reporting 100% utilization gets a small weight but stays in priority 0.
- **Failover targets start without weights.** Remote endpoints get no traffic before failover, so they have sent no reports. They start at equal weight and become weighted after `weightStabilizationPeriod`.
- **Weights are computed per priority.** The default (median) weight is computed within each priority, so idle lower-priority endpoints show `weight: 1` in `istioctl pc ep` even while priority 0 shows real weights (observed: `orca-busy` 1000 at priority 0, `orca-idle` 1 at priority 1).
- **Weighting stays within a priority.** CSWRR never moves traffic from a busy local endpoint to an idle remote one.

## Making sure failover happens

Failover is Envoy's priority logic, not part of CSWRR. Traffic moves to priority 1 only when **priority 0 doesn't have enough healthy hosts**.
CSWRR only spreads traffic inside the priority Envoy picked.

### When Envoy moves traffic

```
P0 load = min(100%, healthy_P0_hosts / total_P0_hosts × 140%)
```

The rest goes to priority 1. istiod doesn't override the overprovisioning factor, so Envoy's default of 140 applies, and failover is gradual:

| Healthy in P0 | Traffic to P0 | Traffic to P1 |
| ------------- | ------------- | ------------- |
| 100%          | 100%          | 0%            |
| 75%           | 100%          | 0%            |
| 50%           | 70%           | 30%           |
| 0%            | 0%            | 100%          |

Failover only starts once fewer than about 72% of P0 hosts are healthy. With one host per priority it is all or nothing.

### Checklist

1. **Endpoints are split into priorities.** `/clusters` must show both `::priority::0` and `::priority::1`.
   If everything is priority 0, there is nothing to fail over to. This happens when every endpoint matches the client on the `failoverPriority` labels or zone/region, or when nodes have no topology labels.
1. **`outlierDetection` is set.** Without it istiod doesn't assign priorities, unless failover is forced by `trafficDistribution: PreferClose`.
1. **Hosts can actually be ejected.**
   - `maxEjectionPercent` must allow it. Envoy ejects only if `(ejected + 1) / hosts ≤ maxEjectionPercent`, default 10%, so with fewer than 10 hosts nothing is ever ejected. Use `100` if a whole priority should be able to drain.
   - The errors must be ones outlier detection counts. `consecutive5xxErrors` counts 5xx; `consecutiveGatewayErrors` counts only 502/503/504; connection failures count too. 4xx and slow responses don't.
   - An ejected host comes back after `baseEjectionTime × number of times ejected`.
1. **Kubernetes health counts too.** A pod failing its readiness probe drops out of EDS, which also lowers P0 health without any outlier detection.
1. **Panic mode doesn't get in the way.** Istio sets `healthyPanicThreshold: {}` (0), which turns panic mode off, so Envoy won't spread traffic across unhealthy P0 hosts.

Things that **do not** trigger failover:

- **High ORCA utilization.** A P0 host reporting `application_utilization=1.0` is still healthy. It gets a small weight, but traffic stays in P0.
- **Latency.** Istio's outlier detection doesn't eject slow hosts.

### Seeing ejections

Outlier stats are filtered out of the Istio proxy stats by default. Add this annotation to the client pod template to include them:

```yaml
metadata:
  annotations:
    proxy.istio.io/config: |
      proxyStatsMatcher:
        inclusionRegexps: [".*outlier_detection.*"]
```

```shell
kubectl exec deploy/sleep -c istio-proxy -- pilot-agent request GET stats | grep orca | grep -E "ejections_(active|enforced_total|overflow):"
kubectl exec deploy/sleep -c istio-proxy -- pilot-agent request GET clusters | grep orca | grep -E "priority|health_flags"
```

- `ejections_enforced_total`: hosts actually ejected. Ejected hosts show `health_flags::/failed_outlier_check`.
- `ejections_overflow`: ejections blocked by `maxEjectionPercent`. If this grows while requests keep failing, raise `maxEjectionPercent`.

### Demo: gradual failover

Add a second P0 host that always returns 503, so P0 has 2 hosts and 1 of them fails:

| Pod             | `version` | Priority (client has `version: busy`) | Behaviour  |
| --------------- | --------- | -------------------------------------- | ---------- |
| `orca-busy`     | busy      | 0                                      | 200        |
| `orca-busy-503` | busy      | 0                                      | always 503 |
| `orca-idle`     | idle      | 1                                      | 200        |

```shell
kubectl create configmap orca-nginx-503 --from-file=default.conf.template=/dev/stdin <<'EOF'
server {
  listen 8080;
  location / {
    return 503 "${POD_NAME} 503\n";
  }
}
EOF
# copy of orca-busy with labels version=busy,broken=true and configMap orca-nginx-503
kubectl patch deploy sleep --type merge -p '{"spec":{"template":{"metadata":{"labels":{"version":"busy"}}}}}'
```

Use the failover `DestinationRule` above with `failoverPriority: [version]`, then send traffic:

```shell
kubectl exec deploy/sleep -- sh -c 'for i in $(seq 1 200); do curl -s orca:8080; done' | awk '{print $1}' | sort | uniq -c
```

Observed results:

| `maxEjectionPercent` | Result                                                                                                                              |
| -------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| default (10%)        | 100 `orca-busy`, 100 × 503 from `orca-busy-503`, 0 `orca-idle`. `ejections_overflow: 120`, `ejections_enforced_total: 0`. No failover. |
| `100`                | 1 × 503, then 138 `orca-busy` / 61 `orca-idle` (next run 141 / 59). `ejections_enforced_total: 1`. The 70/30 split from the table above. |

## Limitations

- **Each client decides on its own.** Weights are learned per client sidecar from its own responses; there is no shared view, so low-traffic clients learn slowly.
- **Per-request reports only.** ORCA's out-of-band streaming mode is not used, so an endpoint that gets no requests sends no reports and its weight expires back to the default.
- **Backends are trusted.** A backend that always reports low utilization will attract more traffic.
- **No session affinity.** It can't be combined with `consistentHash`.
