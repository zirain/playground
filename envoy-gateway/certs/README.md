# Envoy Gateway Certificates

Run `./gen-cert.sh` to (re)generate all cert material below and sync the
matching Secrets/ConfigMaps into the cluster. It's idempotent — re-run it any
time a cert expires or a demo needs to be reset.

| CA (self-signed)  | Leaf                        | Secret               | ConfigMap (`ca.crt` key) | Used by                            |
| ------------------ | --------------------------- | -------------------- | ------------------------ | ----------------------------------- |
| `example.com`       | `www.example.com`           | `example-cert`        | `example-ca`              | `httproute/backend-tls.yaml`        |
| `example.com`       | `tls-backend-1.example.com` | `tls-backend-1-cert`  | `example-ca`              | `httproute/mirror/tls-backend.yaml` |
| `example.org`       | `www.example.org`           | `example-org-cert`    | `example-org-ca`          | -                                    |
| `example.org`       | `mirror.example.com`        | `mirror-tls-backend-cert` | `example-org-ca`     | `httproute/mirror/tls-backend.yaml` |
| `clientca`          | `client`                    | `example-client-cert` (in `envoy-gateway-system`) | `example-client-ca` | `EnvoyProxy/backend-mtls/*.yaml` |

`tls-backend-1` and the mirror target intentionally sit under different base
domains and CAs, so the mirror demo mirrors a request to a backend with a
genuinely different certificate.
