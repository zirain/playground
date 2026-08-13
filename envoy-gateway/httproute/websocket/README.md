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