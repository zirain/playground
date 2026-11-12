#!/usr/bin/env bash
# Usage: traffic.sh [N] [kubeconfig dir]  - send N requests to whereami via both TEG gateways and the mesh client, count per cluster/pod
DEF="$(cd "$(dirname "$0")/../../../istio/multicluster" && pwd)/.kube"
N=${1:-60}; K=${2:-$DEF}
gw() { kubectl --kubeconfig $K/primary -n envoy-gateway-system get svc -l gateway.envoyproxy.io/owning-gateway-name=$1 -o jsonpath='{.items[0].metadata.name}'; }
kubectl --kubeconfig $K/primary -n envoy-gateway-system port-forward svc/$(gw sidecar-endpoint) 18190:80 >/dev/null 2>&1 & P1=$!
kubectl --kubeconfig $K/primary -n envoy-gateway-system port-forward svc/$(gw sidecar-service) 18191:80 >/dev/null 2>&1 & P2=$!
trap 'kill $P1 $P2' EXIT; sleep 3
for g in "routingType=Endpoint:18190" "routingType=Service:18191"; do echo "== TEG ${g%%:*}"
  for i in $(seq $N); do r=$(curl -s -w ' [%{http_code}]' localhost:${g##*:}/ -H 'Host: whereami.example.com'); echo "${r//$'\n'/}"; done | sort | uniq -c; done
echo "== mesh client (sidecar, no TEG)"
kubectl --kubeconfig $K/primary -n h2test exec meshclient -c meshclient -- sh -c "for i in \$(seq $N); do curl -s http://whereami:8000/; done" | sort | uniq -c
