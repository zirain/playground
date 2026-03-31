# Envoy Gateway Merge Gateways Demo

This demo demonstrates the merge gateways feature in Envoy Gateway, which allows multiple Gateway resources to share the same underlying infrastructure (load balancer).

## Overview

The merge gateways feature is useful when:
- Multiple teams want to share the same load balancer
- You want to reduce infrastructure costs by consolidating gateways
- You need to manage different hostnames/domains through separate Gateway resources

## Architecture

This demo creates:
1. **EnvoyProxy** resource with `mergeGateways: true` enabled
2. **Three Gateway resources**:
   - `gateway-v1`: Handles traffic for `v1.example.com`
   - `gateway-v2`: Handles traffic for `v2.example.com`
   - `gateway-v3`: Handles traffic for `v3.example.com`
3. **Three HTTPRoutes**: One for each gateway, routing to `infra-backend-v1/v2/v3` respectively
4. **Backend services**: Uses existing `infra-backend-v1`, `infra-backend-v2`, and `infra-backend-v3` services
5. **EnvoyExtensionPolicy**: Attaches a Lua filter to the `GatewayClass` and individual `Gateway`.
6. **EnvoyPatchPolicy**: Demonstrates low-level xDS patching targeting the `GatewayClass` (useful for features not yet in the official API).

All three gateways will be **merged into a single Envoy proxy deployment** sharing the same load balancer.

## Prerequisites

- Kubernetes cluster
- Envoy Gateway installed
- `infra-backend-v1`, `infra-backend-v2`, `infra-backend-v3` services deployed (from e2e-manifests or create them separately)

## Installation

1. Apply the manifest:
```bash
kubectl apply -f merge-gateway.yaml
```

2. Wait for all resources to be ready:
```bash
kubectl wait --timeout=5m -n envoy-gateway-system deployment/envoy-merge-proxy --for=condition=Available
kubectl wait --timeout=5m gateway/gateway-v1 --for=condition=Programmed
kubectl wait --timeout=5m gateway/gateway-v2 --for=condition=Programmed
kubectl wait --timeout=5m gateway/gateway-v3 --for=condition=Programmed
```

3. Get the Gateway address:
```bash
export GATEWAY_HOST=$(kubectl get gateway gateway-v1 -o jsonpath='{.status.addresses[0].value}')
echo "Gateway Host: $GATEWAY_HOST"
```

## Testing

### Test v1 Route
```bash
curl -v -H "Host: v1.example.com" http://$GATEWAY_HOST/
```

Expected output should show:
- Pod name from `infra-backend-v1` deployment.
- Response header `X-Lua-Response: true` (from GatewayClass policy).
- In Envoy logs: `V1-specific Lua filter executed`.

### Test v2 Route
```bash
curl -v -H "Host: v2.example.com" http://$GATEWAY_HOST/
```

Expected output should show:
- Pod name from `infra-backend-v2` deployment.
- Response header `X-Lua-Response: true` (from GatewayClass policy).
- It should **NOT** have v1-specific headers/logs.

### Test v3 Route
```bash
curl -H "Host: v3.example.com" http://$GATEWAY_HOST/
```

Expected output should show pod name from `infra-backend-v3` deployment.

## Verification

### Verify Gateway Merge
Check that all gateways point to the same service:

```bash
# All three gateways should have the same external IP/hostname
kubectl get gateway gateway-v1 -o jsonpath='{.status.addresses[0].value}'
kubectl get gateway gateway-v2 -o jsonpath='{.status.addresses[0].value}'
kubectl get gateway gateway-v3 -o jsonpath='{.status.addresses[0].value}'
```

### Verify Single Envoy Deployment
Check that only one Envoy proxy deployment was created:

```bash
kubectl get deployment -n envoy-gateway-system | grep envoy
```

You should see only **one deployment** instead of three separate ones.

### Check Envoy Configuration
```bash
# Get the proxy pod name
PROXY_POD=$(kubectl get pods -n envoy-gateway-system -l gateway.envoyproxy.io/owning-gatewayclass=eg-merge-demo -o jsonpath='{.items[0].metadata.name}')

# Dump the Envoy config and verify the patch (e.g., look for per_connection_buffer_limit_bytes)
kubectl exec -n envoy-gateway-system $PROXY_POD -- curl -s http://localhost:19000/config_dump | grep per_connection_buffer_limit_bytes
```

### Check Listeners
The merged Envoy should have all listeners configured:

```bash
kubectl exec -n envoy-gateway-system $PROXY_POD -- curl -s http://localhost:19000/config_dump | grep -A2 '"name": "default'
```

## Key Concepts

### MergeGateways Feature
- Enabled via `spec.mergeGateways: true` in the EnvoyProxy resource
- All Gateways using the same GatewayClass will share the same Envoy proxy deployment
- Each Gateway's listeners are merged into a single Envoy configuration
- Reduces infrastructure costs and management overhead

### Benefits
1. **Cost Reduction**: Single load balancer instead of multiple
2. **Simplified Management**: One proxy to monitor and manage
3. **Resource Efficiency**: Shared infrastructure reduces resource usage
4. **Flexibility**: Teams can still manage their own Gateway resources independently

### Limitations
- All merged gateways must use the same GatewayClass
- Listener conflicts (same port + hostname) must be avoided
- All merged gateways share the same scaling characteristics

## Cleanup

```bash
kubectl delete -f merge-gateway.yaml
```

## Troubleshooting

### Gateways not merging
Check the EnvoyProxy resource has `mergeGateways: true`:
```bash
kubectl get envoyproxy -n envoy-gateway-system merge-proxy -o yaml
```

### Routes not working
Check HTTPRoute status:
```bash
kubectl get httproute route-v1 -o yaml
kubectl describe httproute route-v1
```

### Backend services not found
Make sure the infra-backend services are deployed:
```bash
kubectl get svc infra-backend-v1 infra-backend-v2 infra-backend-v3
```

### Check Envoy Gateway logs
```bash
kubectl logs -n envoy-gateway-system deployment/envoy-gateway -f
```

### Check Envoy Proxy logs
```bash
kubectl logs -n envoy-gateway-system $PROXY_POD -f
```

## References

- [Envoy Gateway Merge Gateways Documentation](https://gateway.envoyproxy.io/latest/user/merge-gateways/)
- [Gateway API Specification](https://gateway-api.sigs.k8s.io/)
