# Istio canary upgrade: 503s during workload rollout

Upgrading the data plane with revisions (1-26-8 → 1-28-4 → 1-31-1) while sending
traffic through the ingress gateway gives 1–2 `503`s every time a workload pod is
replaced. This page covers the setup, the cause, and what does and doesn't fix it.

**In short:** the 503s come from the old pod's app exiting before the gateway stops
sending to it. A backend that keeps serving for a few seconds after SIGTERM (a shutdown
delay) had no errors in our runs, while httpbin, which exits right away, failed a
request each time ([Verification](#verification)). For apps without a shutdown delay,
a `preStop` sleep has the same effect.

## Setup

| File | Purpose |
|---|---|
| [iop.yaml](iop.yaml) | Control plane install (`minimal` profile, access logs on) |
| [gateway.yaml](gateway.yaml) | Ingress gateway Deployment, Service and RBAC; `upgrade.sh` sets `istio.io/rev` |
| [expose-httpbin.yaml](expose-httpbin.yaml) | Gateway + VirtualService: `/agnhost/*` to agnhost, everything else to httpbin |
| [agnhost.yaml](agnhost.yaml) | Backend with a 5s shutdown delay, for comparison with httpbin |
| [upgrade.sh](upgrade.sh) | Runs the upgrade with vegeta traffic in the background |

`upgrade.sh` does the following:

1. Starts `vegeta attack` against the gateway in the background.
2. Applies `gateway.yaml` with `istio.io/rev=${REVISION}` and waits for the rollout.
3. Relabels the namespace with `istio.io/rev=${REVISION}` and removes `istio-injection`,
   which would otherwise take precedence.
4. Restarts every deployment in the namespace and waits for each one.
5. Runs `istioctl proxy-status`, then prints the vegeta report.

```sh
REVISION=1-31-1 DURATION=60s ./upgrade.sh

# Compare both backends in the same rollout
(echo "GET http://172.18.255.201/agnhost/hostname" | vegeta attack -rate=10 -duration=60s > agnhost-results.bin) &
REVISION=1-31-1 DURATION=60s RESULTS=httpbin-results.bin ./upgrade.sh
wait; vegeta report < agnhost-results.bin
```

Make `DURATION` longer than the whole upgrade, otherwise the last restarts send no traffic.

## What we saw

Every run had one or two failed requests, always to the **old** httpbin pod, within
about a second of it being told to stop. Two forms showed up in the gateway access log:

```
"GET /get HTTP/1.1" 503 UC upstream_reset_before_response_started{connection_termination} ... "10.244.1.120:8080"
"GET /get HTTP/1.1" 503 - via_upstream - "-" 0 190 52 52 ... "10.244.1.127:8080"
```

- `503 UC`: the old pod closed the gateway's keep-alive connection while a request was on it.
- `503 via_upstream` (190-byte body, about 50ms upstream time): the old pod's sidecar
  accepted the request but couldn't reach the app, so it replied with its own 503.

## Why it happens

It's a race between the old pod shutting down and the gateway learning that it's gone.

1. **Old pod deleted.** The rollout deletes it once the new pod is Ready.
2. **App exits.** `istio-proxy` runs as a native sidecar (an init container), so the
   app container gets SIGTERM first. go-httpbin exits almost immediately.
3. **Sidecar still accepts requests.** It stays up until the app has exited, but it
   has nowhere to forward them, so it returns `503 via_upstream`. Slightly later in the
   shutdown, connections are closed, which gives `503 UC`.
4. **Gateway catches up.** The endpoint removal goes through istiod to the gateway in
   about 100ms. Requests sent in that window go to the old pod.
5. **No retry.** The default retry policy is
   `connect-failure,refused-stream,unavailable,cancelled,retriable-status-codes`, with an
   empty status-code list. It covers neither `reset` nor an upstream `503`.

Running a single replica makes this worse, since all traffic goes to the pod being replaced.

## What doesn't fix it

| Change | Why not |
|---|---|
| Rollout strategy | With 1 replica, the defaults (`maxSurge: 25%` → 1, `maxUnavailable: 25%` → 0) already start the new pod first. The strategy decides *when* the old pod is deleted, not what happens while it shuts down. `minReadySeconds` and more replicas only shrink the impact. |
| Readiness probe | Controls when a pod *starts* getting traffic. A terminating pod is removed from endpoints straight away whatever its probe says. Still worth adding, so the old pod isn't killed before the new app is listening. |
| `PILOT_AUTO_SEND_UNHEALTHY_ENDPOINTS: "false"` | Only decides whether not-ready endpoints are sent to proxies. Terminating pods are dropped either way, and the delay is the push to the gateway. |
| Usual graceful shutdown (drain in-flight requests, close listener right away) | The failing requests are *new* ones arriving after SIGTERM, which are still refused. |

## What fixes it

Both of the first two fixes keep the app serving for a few seconds after SIGTERM, so the
gateway switches away before the app stops. The native sidecar waits for the app to exit
either way.

**Sizing the delay.** The shutdown delay (or `preStop` sleep) must be longer than the
time it takes istiod to push the endpoint removal to the gateway over xDS. Size it from
istiod's metrics:

| Metric | What it measures |
|---|---|
| `pilot_proxy_convergence_time` | Config change until a proxy has received all its config. Use this one. |
| `pilot_debounce_time` | Time a change waits in debouncing before the push is queued (`PILOT_DEBOUNCE_AFTER`, default 100ms) |
| `pilot_proxy_queue_time` | Time a proxy waits in the push queue; grows when istiod is busy |

```promql
histogram_quantile(0.99, sum(rate(pilot_proxy_convergence_time_bucket[5m])) by (le))
```

These start once istiod has seen the change. They don't include the time for the pod's
removal to travel from Kubernetes (EndpointSlice) to istiod. On this cluster every push
converged in under 0.1s by the metric (about 1ms on average), while the access logs showed
about 100ms between SIGTERM and the gateway switching pods. So take the p99, add generous
margin, and prefer several seconds over fine-tuning: 5s is well clear here, and the extra
wait only slows pod termination.

1. **Use a backend that supports a shutdown delay.** Many servers and frameworks have a
   setting that keeps serving for a while after SIGTERM before draining. agnhost's
   `--delay-shutdown=5` is the one verified below. This keeps the behaviour with the app
   and needs no Kubernetes config.
2. **Add a `preStop` sleep** for apps that can't be changed, like httpbin, which exits
   right away. go-httpbin is distroless, so use the `sleep` action (Kubernetes 1.30+),
   not `exec sleep`:
   ```yaml
   lifecycle:
     preStop:
       sleep:
         seconds: 5
   ```
3. **Retries as a safety net**, on the VirtualService. Safe here because the test only
   sends GETs:
   ```yaml
   retries:
     attempts: 2
     retryOn: connect-failure,refused-stream,unavailable,cancelled,reset,503
   ```
   The default `previous_hosts` predicate sends a retry to a different pod.

## Verification

agnhost (`netexec --delay-shutdown=5`) and httpbin ran side by side, with both old pods
told to stop in the same second and retries off. Each run was 10 rps for 60s per backend.

| Upgrade step | httpbin (exits right away) | agnhost (5s shutdown delay) |
|---|---|---|
| 1-26-8 → 1-28-4 | 599 × 200, **1 × 503** | **600 × 200** |
| 1-28-4 → 1-31-1 | 599 × 200, **1 × 503** | **600 × 200** |

In the first run, agnhost's old pod served six more requests after SIGTERM, all 200,
before traffic moved to the new pod. Restarting the gateway caused no errors on either
backend.

Not yet verified on the cluster: the `preStop` sleep on httpbin and the retry policy.
Both are expected to give zero errors; a run with each would confirm it.
