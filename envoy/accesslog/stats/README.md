# Envoy Stats Access Logger

Demonstrates the [`envoy.access_loggers.stats`](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/access_loggers/stats/v3/stats.proto) extension, which emits custom Envoy stats on each access log event instead of writing log lines.

> **Note:** This extension is work-in-progress and not recommended for production use.

## How it works

`listener_0` (port 10000) is configured with a stats access logger (`stat_prefix: ingress_http_al`) that records:

| Stat | Type | Unit | Tags |
|------|------|------|------|
| `ingress_http_al.request_count` | Counter | — | `response_code`, `cluster` |
| `ingress_http_al.response_duration_ms` | Histogram | Milliseconds | `response_code` |
| `ingress_http_al.response_bytes` | Histogram | Bytes | `response_code` |

Routes on `listener_0`:
- `GET /cluster1` → proxied to `cluster1` (backed by `listener_1` on port 10001)
- `GET /` → direct `200` response with `"example body"`

`listener_1` (port 10001) acts as a simple upstream returning `"reply from listener_1"`.

## Run

```bash
docker compose up
```

## Test

```bash
# Direct response
curl http://localhost:10000/

# Proxied to cluster1
curl http://localhost:10000/cluster1
```

## Inspect stats

```bash
curl http://localhost:9901/stats | grep ingress_http_al
```
