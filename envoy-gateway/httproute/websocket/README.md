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