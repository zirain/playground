# Envoy Gateway WebSocket demo

This demo routes WebSocket upgrade requests at `/echo` through Envoy Gateway to a WebSocket echo server. Envoy Gateway automatically configures the WebSocket upgrade for an `HTTPRoute`; no `EnvoyPatchPolicy` is needed.

## Prerequisites

- A Kubernetes cluster with Envoy Gateway installed and a `GatewayClass` named `eg`.
- `kubectl` and [websocat](https://github.com/vi/websocat).

## Run

Apply the demo and wait for its route and backend:

```sh
kubectl apply -f websocket.yaml
kubectl wait --for=condition=programmed gateway/websocket-gateway -n websocket-demo --timeout=2m
kubectl rollout status deployment/websocket-echo -n websocket-demo
```

Retrieve the Gateway address, then open a WebSocket connection:

```sh
GATEWAY_HOST=$(kubectl get gateway websocket-gateway -n websocket-demo \
  -o jsonpath='{.status.addresses[0].value}')
websocat "ws://${GATEWAY_HOST}/echo"
```

Type a message and the echo server returns it. A successful handshake responds with HTTP `101 Switching Protocols`.

For clusters without an externally reachable Gateway address, port-forward the generated Envoy Service and use the local address:

```sh
GATEWAY_SERVICE=$(kubectl get service -n websocket-demo \
  -l gateway.envoyproxy.io/owning-gateway-name=websocket-gateway \
  -o jsonpath='{.items[0].metadata.name}')
kubectl port-forward -n websocket-demo "service/${GATEWAY_SERVICE}" 8080:80
websocat ws://127.0.0.1:8080/echo
```

## Inspect the generated route

The generated Envoy route should contain `upgrade_type: websocket`:

```sh
ENVOY_POD=$(kubectl get pods -n websocket-demo \
  -l gateway.envoyproxy.io/owning-gateway-name=websocket-gateway \
  -o jsonpath='{.items[0].metadata.name}')
kubectl exec -n websocket-demo "${ENVOY_POD}" -c envoy -- \
  curl -s 127.0.0.1:19000/config_dump | grep -A1 -B4 websocket
```

## Clean up

```sh
kubectl delete -f websocket.yaml
```

## SSL backend variant

`websocket-ssl-backend.yaml` is the same demo, but Envoy Gateway originates TLS to
the backend (an "SSL backend"). The client-facing listener is still plain
`HTTP`/`ws://`; only the Gateway-to-backend hop is encrypted. Since
`jmalloc/echo-server` has no TLS support, the pod adds an `nginx` sidecar that
terminates TLS on `8443` and reverse-proxies the WebSocket upgrade to the echo
server on `127.0.0.1:8080`. A `BackendTLSPolicy` tells Envoy Gateway to connect
to that port over TLS and validate the backend's certificate.

### Generate a backend certificate

The Secret/ConfigMap it depends on aren't in the manifest (this repo generates
certs out-of-band, see [`certs/README.md`](../../certs/README.md)). Create a
self-signed cert/key for `websocket-backend.example.com` — SAN is required,
CN alone isn't enough for most TLS clients including Envoy:

```sh
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 \
  -subj '/CN=websocket-backend.example.com/O=example organization' \
  -addext 'subjectAltName=DNS:websocket-backend.example.com' \
  -keyout websocket-backend.key -out websocket-backend.crt

kubectl create namespace websocket-ssl-demo
kubectl create secret tls websocket-backend-tls -n websocket-ssl-demo \
  --cert=websocket-backend.crt --key=websocket-backend.key
# Self-signed, so the leaf cert is also its own trust anchor.
kubectl create configmap websocket-backend-ca -n websocket-ssl-demo \
  --from-file=ca.crt=websocket-backend.crt
```

### Run

```sh
kubectl apply -f websocket-ssl-backend.yaml
kubectl wait --for=condition=programmed gateway/websocket-ssl-gateway -n websocket-ssl-demo --timeout=2m
kubectl rollout status deployment/websocket-echo-tls -n websocket-ssl-demo
```

The client side is unchanged — connect with `ws://`, same as the plain demo:

```sh
GATEWAY_HOST=$(kubectl get gateway websocket-ssl-gateway -n websocket-ssl-demo \
  -o jsonpath='{.status.addresses[0].value}')
websocat "ws://${GATEWAY_HOST}/echo"
```

To confirm Envoy is actually using TLS upstream, check the cluster's transport
socket in the config dump:

```sh
ENVOY_POD=$(kubectl get pods -n websocket-ssl-demo \
  -l gateway.envoyproxy.io/owning-gateway-name=websocket-ssl-gateway \
  -o jsonpath='{.items[0].metadata.name}')
kubectl exec -n websocket-ssl-demo "${ENVOY_POD}" -c envoy -- \
  curl -s 127.0.0.1:19000/config_dump | grep -A5 '"name": ".*websocket-echo-tls'
```

### Clean up

```sh
kubectl delete -f websocket-ssl-backend.yaml
kubectl delete secret websocket-backend-tls -n websocket-ssl-demo
kubectl delete configmap websocket-backend-ca -n websocket-ssl-demo
```

## SSL frontend + SSL backend variant

`websocket-wss-ssl-backend.yaml` takes the SSL-backend demo further: the
client-facing listener now terminates TLS too (`HTTPS`/port `443`), so clients
connect with `wss://` instead of `ws://`. It uses a separate certificate
(`websocket-frontend-tls`) from the one the backend TLS hop uses
(`websocket-backend-tls`) — TLS is terminated independently on each hop. It
runs in its own `websocket-wss-ssl-demo` namespace, so it can coexist with the
other demos.

### Generate certificates

Two unrelated self-signed certs are needed: one for the Gateway's
client-facing listener, one for the backend (same as the SSL-backend variant):

```sh
# Frontend cert — terminated by the Gateway listener itself.
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 \
  -subj '/CN=websocket-frontend.example.com/O=example organization' \
  -addext 'subjectAltName=DNS:websocket-frontend.example.com' \
  -keyout websocket-frontend.key -out websocket-frontend.crt

# Backend cert — terminated by the nginx sidecar, validated via BackendTLSPolicy.
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 \
  -subj '/CN=websocket-backend.example.com/O=example organization' \
  -addext 'subjectAltName=DNS:websocket-backend.example.com' \
  -keyout websocket-backend.key -out websocket-backend.crt

kubectl create namespace websocket-wss-ssl-demo
kubectl create secret tls websocket-frontend-tls -n websocket-wss-ssl-demo \
  --cert=websocket-frontend.crt --key=websocket-frontend.key
kubectl create secret tls websocket-backend-tls -n websocket-wss-ssl-demo \
  --cert=websocket-backend.crt --key=websocket-backend.key
# Self-signed, so the backend leaf cert is also its own trust anchor.
kubectl create configmap websocket-backend-ca -n websocket-wss-ssl-demo \
  --from-file=ca.crt=websocket-backend.crt
```

### Run

```sh
kubectl apply -f websocket-wss-ssl-backend.yaml
kubectl wait --for=condition=programmed gateway/websocket-wss-ssl-gateway -n websocket-wss-ssl-demo --timeout=2m
kubectl rollout status deployment/websocket-echo-tls -n websocket-wss-ssl-demo
```

Connect with `wss://`. The frontend cert is self-signed, so `websocat` needs
`-k`/`--insecure` to skip certificate validation:

```sh
GATEWAY_HOST=$(kubectl get gateway websocket-wss-ssl-gateway -n websocket-wss-ssl-demo \
  -o jsonpath='{.status.addresses[0].value}')
websocat -k "wss://${GATEWAY_HOST}/echo"
```

To confirm both hops are actually using TLS, check the listener's transport
socket (client-facing) and the cluster's transport socket (backend-facing) in
the config dump:

```sh
ENVOY_POD=$(kubectl get pods -n websocket-wss-ssl-demo \
  -l gateway.envoyproxy.io/owning-gateway-name=websocket-wss-ssl-gateway \
  -o jsonpath='{.items[0].metadata.name}')
kubectl exec -n websocket-wss-ssl-demo "${ENVOY_POD}" -c envoy -- \
  curl -s 127.0.0.1:19000/config_dump | grep -A5 '"name": "wss"'
kubectl exec -n websocket-wss-ssl-demo "${ENVOY_POD}" -c envoy -- \
  curl -s 127.0.0.1:19000/config_dump | grep -A5 '"name": ".*websocket-echo-tls'
```

### Clean up

```sh
kubectl delete -f websocket-wss-ssl-backend.yaml
kubectl delete secret websocket-frontend-tls websocket-backend-tls -n websocket-wss-ssl-demo
kubectl delete configmap websocket-backend-ca -n websocket-wss-ssl-demo
```

## Broken SSL backend variant

`websocket-ssl-backend-500.yaml` is the SSL-backend demo again, but the backend
always answers `500`. It's the same `nginx` TLS termination on `8443`, minus the
`proxy_pass` to an echo server — `location /` just does `return 500;`. It models
a backend that's reachable and completes the TLS handshake but is broken at the
HTTP layer, for exercising Envoy Gateway's error handling on an SSL backend
rather than the happy path. It runs in its own `websocket-ssl-500-demo`
namespace, so it can coexist with the other two demos.

### Generate a backend certificate

```sh
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 \
  -subj '/CN=websocket-backend-500.example.com/O=example organization' \
  -addext 'subjectAltName=DNS:websocket-backend-500.example.com' \
  -keyout websocket-backend-500.key -out websocket-backend-500.crt

kubectl create namespace websocket-ssl-500-demo
kubectl create secret tls websocket-backend-500-tls -n websocket-ssl-500-demo \
  --cert=websocket-backend-500.crt --key=websocket-backend-500.key
# Self-signed, so the leaf cert is also its own trust anchor.
kubectl create configmap websocket-backend-500-ca -n websocket-ssl-500-demo \
  --from-file=ca.crt=websocket-backend-500.crt
```

### Run

```sh
kubectl apply -f websocket-ssl-backend-500.yaml
kubectl wait --for=condition=programmed gateway/websocket-ssl-500-gateway -n websocket-ssl-500-demo --timeout=2m
kubectl rollout status deployment/websocket-tls-500 -n websocket-ssl-500-demo
```

Connect with `ws://`, same as the other demos. Since the backend always returns
`500`, the handshake fails instead of upgrading (no `101 Switching Protocols`):

```sh
GATEWAY_HOST=$(kubectl get gateway websocket-ssl-500-gateway -n websocket-ssl-500-demo \
  -o jsonpath='{.status.addresses[0].value}')
websocat "ws://${GATEWAY_HOST}/echo"
# websocat: WebSocketError: Received unexpected status code (500 Internal Server Error)
```

To confirm the `500` is coming from the backend and not from a failed TLS
handshake, check the cluster's transport socket in the config dump:

```sh
ENVOY_POD=$(kubectl get pods -n websocket-ssl-500-demo \
  -l gateway.envoyproxy.io/owning-gateway-name=websocket-ssl-500-gateway \
  -o jsonpath='{.items[0].metadata.name}')
kubectl exec -n websocket-ssl-500-demo "${ENVOY_POD}" -c envoy -- \
  curl -s 127.0.0.1:19000/config_dump | grep -A5 '"name": ".*websocket-tls-500'
```

### Clean up

```sh
kubectl delete -f websocket-ssl-backend-500.yaml
kubectl delete secret websocket-backend-500-tls -n websocket-ssl-500-demo
kubectl delete configmap websocket-backend-500-ca -n websocket-ssl-500-demo
```

## Certificate mismatch variants

The next two demos are also SSL-backend demos with a healthy echo server behind
the `tls-proxy` sidecar, but each has a deliberately broken `BackendTLSPolicy` so
the TLS handshake to the backend fails before any WebSocket traffic gets
through. They isolate the two distinct ways an upstream cert check can fail:
the hostname doesn't match the cert, or the cert doesn't chain to the trusted
CA. Both should make Envoy Gateway reject the upstream connection and return
`503` — never the `500`/`101` the app itself would produce.

### Hostname mismatch

`websocket-ssl-backend-hostname-mismatch.yaml` presents a cert that's valid and
self-trusted (the `BackendTLSPolicy`'s CA ref is the cert's own issuer), but
the policy's `validation.hostname` doesn't match the cert's SAN — as if the
cert were rotated to a new domain and the policy never got updated. Runs in
its own `websocket-hostname-mismatch-demo` namespace.

#### Generate a backend certificate

```sh
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 \
  -subj '/CN=websocket-backend-hostname-mismatch.example.com/O=example organization' \
  -addext 'subjectAltName=DNS:websocket-backend-hostname-mismatch.example.com' \
  -keyout websocket-backend-hostname-mismatch.key -out websocket-backend-hostname-mismatch.crt

kubectl create namespace websocket-hostname-mismatch-demo
kubectl create secret tls websocket-backend-tls -n websocket-hostname-mismatch-demo \
  --cert=websocket-backend-hostname-mismatch.crt --key=websocket-backend-hostname-mismatch.key
# Self-signed, so the leaf cert is also its own trust anchor — the CA ref is
# correct. It's the hostname in the BackendTLSPolicy (websocket-backend.example.com)
# that doesn't match this cert's SAN.
kubectl create configmap websocket-backend-ca -n websocket-hostname-mismatch-demo \
  --from-file=ca.crt=websocket-backend-hostname-mismatch.crt
```

#### Run

```sh
kubectl apply -f websocket-ssl-backend-hostname-mismatch.yaml
kubectl wait --for=condition=programmed gateway/websocket-hostname-mismatch-gateway -n websocket-hostname-mismatch-demo --timeout=2m
kubectl rollout status deployment/websocket-echo-tls -n websocket-hostname-mismatch-demo
```

Connect with `ws://`, same as the other demos. The upstream TLS handshake
fails hostname verification, so Envoy Gateway never reaches the (healthy)
echo server:

```sh
GATEWAY_HOST=$(kubectl get gateway websocket-hostname-mismatch-gateway -n websocket-hostname-mismatch-demo \
  -o jsonpath='{.status.addresses[0].value}')
websocat "ws://${GATEWAY_HOST}/echo"
# websocat: WebSocketError: Received unexpected status code (503 Service Unavailable)
```

To confirm the failure is a cert verification error and not the app rejecting
the request, check Envoy's SSL stats for the backend cluster:

```sh
ENVOY_POD=$(kubectl get pods -n websocket-hostname-mismatch-demo \
  -l gateway.envoyproxy.io/owning-gateway-name=websocket-hostname-mismatch-gateway \
  -o jsonpath='{.items[0].metadata.name}')
kubectl exec -n websocket-hostname-mismatch-demo "${ENVOY_POD}" -c envoy -- \
  curl -s 127.0.0.1:19000/stats | grep -i 'ssl.*fail'
```

#### Clean up

```sh
kubectl delete -f websocket-ssl-backend-hostname-mismatch.yaml
kubectl delete secret websocket-backend-tls -n websocket-hostname-mismatch-demo
kubectl delete configmap websocket-backend-ca -n websocket-hostname-mismatch-demo
```

### CA mismatch

`websocket-ssl-backend-ca-mismatch.yaml` presents a cert whose hostname
matches the policy fine, but the `websocket-backend-ca` ConfigMap is seeded
with a completely unrelated self-signed cert — not the one the backend
actually presents — so the chain of trust itself fails to verify. Runs in its
own `websocket-ca-mismatch-demo` namespace.

#### Generate the certificates

Two unrelated self-signed certs: the one the backend actually serves, and a
decoy that gets configured as the (wrong) trusted CA:

```sh
# The cert nginx actually presents.
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 \
  -subj '/CN=websocket-backend.example.com/O=example organization' \
  -addext 'subjectAltName=DNS:websocket-backend.example.com' \
  -keyout websocket-backend.key -out websocket-backend.crt

# An unrelated cert, only its public half is used — as the (wrong) trust anchor.
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 \
  -subj '/CN=wrong-ca.example.com/O=example organization' \
  -addext 'subjectAltName=DNS:wrong-ca.example.com' \
  -keyout wrong-ca.key -out wrong-ca.crt

kubectl create namespace websocket-ca-mismatch-demo
kubectl create secret tls websocket-backend-tls -n websocket-ca-mismatch-demo \
  --cert=websocket-backend.crt --key=websocket-backend.key
# Deliberately the wrong cert — doesn't chain to websocket-backend.crt at all.
kubectl create configmap websocket-backend-ca -n websocket-ca-mismatch-demo \
  --from-file=ca.crt=wrong-ca.crt
```

#### Run

```sh
kubectl apply -f websocket-ssl-backend-ca-mismatch.yaml
kubectl wait --for=condition=programmed gateway/websocket-ca-mismatch-gateway -n websocket-ca-mismatch-demo --timeout=2m
kubectl rollout status deployment/websocket-echo-tls -n websocket-ca-mismatch-demo
```

Connect with `ws://`, same as the other demos. The upstream TLS handshake
fails chain-of-trust verification, so Envoy Gateway never reaches the
(healthy) echo server:

```sh
GATEWAY_HOST=$(kubectl get gateway websocket-ca-mismatch-gateway -n websocket-ca-mismatch-demo \
  -o jsonpath='{.status.addresses[0].value}')
websocat "ws://${GATEWAY_HOST}/echo"
# websocat: WebSocketError: Received unexpected status code (503 Service Unavailable)
```

To confirm the failure is a cert verification error and not the app rejecting
the request, check Envoy's SSL stats for the backend cluster:

```sh
ENVOY_POD=$(kubectl get pods -n websocket-ca-mismatch-demo \
  -l gateway.envoyproxy.io/owning-gateway-name=websocket-ca-mismatch-gateway \
  -o jsonpath='{.items[0].metadata.name}')
kubectl exec -n websocket-ca-mismatch-demo "${ENVOY_POD}" -c envoy -- \
  curl -s 127.0.0.1:19000/stats | grep -i 'ssl.*fail'
```

#### Clean up

```sh
kubectl delete -f websocket-ssl-backend-ca-mismatch.yaml
kubectl delete secret websocket-backend-tls -n websocket-ca-mismatch-demo
kubectl delete configmap websocket-backend-ca -n websocket-ca-mismatch-demo
```