# Route to a non-existing cluster

Shows what Envoy does when a route points at a cluster that doesn't exist.

- `validate_clusters: false` on the static `route_config` lets Envoy start even though
  `non-existing-cluster` is not defined. Without it, Envoy refuses the config with
  `route: unknown cluster 'non-existing-cluster'`.
- At request time the router returns `503` with response flag `NC` and response code
  details `cluster_not_found`. The `http.<stat_prefix>.no_cluster` counter goes up.
- `cluster_not_found_response_code` changes the status code (`SERVICE_UNAVAILABLE` by default,
  or `NOT_FOUND` / `INTERNAL_SERVER_ERROR`).

## Weighted clusters

Envoy picks one of the weighted clusters first, then looks that cluster up. Only the
requests that land on the missing cluster fail. The others go through normally.

- With `existing-cluster: 50` and `non-existing-cluster: 50`, 20 requests to `/weighted`
  gave about half 200 and half 503 `NC`. The failed share depends on the missing
  cluster's weight.
- `cluster_not_found_response_code` also applies here. `/weighted-500` sets
  `INTERNAL_SERVER_ERROR`, so its failed share returns 500 instead of 503. The access
  log still shows `NC cluster_not_found`.
- Validation applies to weighted clusters too. With `validate_clusters` on, Envoy
  rejects the config with `route: unknown weighted cluster 'non-existing-cluster'`.

```shell
for i in $(seq 1 20); do curl -s -o /dev/null -w "%{http_code}\n" localhost:10000/weighted; done | sort | uniq -c      # 200 / 503
for i in $(seq 1 20); do curl -s -o /dev/null -w "%{http_code}\n" localhost:10000/weighted-500; done | sort | uniq -c  # 200 / 500
```

## Run

```shell
envoy -c envoy.yaml
```

## Test

```shell
curl -i localhost:10000/                                     # 200, existing-cluster
curl -i localhost:10000/missing                              # 503 NC cluster_not_found
curl -i localhost:10000/missing-404                          # 404 (cluster_not_found_response_code)
curl -i -H 'x-cluster: existing-cluster' localhost:10000/header  # 200
curl -i -H 'x-cluster: nope' localhost:10000/header          # 503 NC
curl -s localhost:15000/stats | grep ingress_http.no_cluster
```

Access log:

```
"GET /missing" 503 NC cluster_not_found cluster=-
"GET /missing-404" 404 NC cluster_not_found cluster=-
```
