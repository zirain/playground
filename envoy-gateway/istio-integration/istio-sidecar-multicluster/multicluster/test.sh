#!/usr/bin/env bash
# Usage: test.sh [kubeconfig dir]  - port protocol matrix; default kubeconfigs from istio/multicluster/.kube
DEF="$(cd "$(dirname "$0")/../../../istio/multicluster" && pwd)/.kube"
K=${1:-$DEF}
B=$(mktemp)
gw() { kubectl --kubeconfig $K/primary -n envoy-gateway-system get svc -l gateway.envoyproxy.io/owning-gateway-name=$1 -o jsonpath='{.items[0].metadata.name}'; }
kubectl --kubeconfig $K/primary -n envoy-gateway-system port-forward svc/$(gw sidecar-endpoint) 18190:80 >/dev/null 2>&1 & P1=$!
kubectl --kubeconfig $K/primary -n envoy-gateway-system port-forward svc/$(gw sidecar-service) 18191:80 >/dev/null 2>&1 & P2=$!
trap 'kill $P1 $P2' EXIT; sleep 3
req() { local c; c=$(curl -s -o $B -w '%{http_code}' localhost:$1$3 -H "Host: $2.example.com"); [[ $2 == nginx* ]] && c="$c $(tr -d '\n' <$B | sed 's/nginx in //')"; echo "$c"; }
printf "%-18s %-30s %-30s %s\n" backend "TEG routingType=Endpoint" "TEG routingType=Service" "mesh client"
for L in primary remote1 remote2; do for s in http-$L:8000:/get http2-$L:8001:/get nginx-h2c-$L:8004:/; do IFS=: read n p u <<<"$s"
  m=$(kubectl --kubeconfig $K/primary -n h2test exec meshclient -c meshclient -- curl -s -o /dev/null -w '%{http_code}' http://$n:$p$u)
  printf "%-18s %-30s %-30s %s\n" $n "$(req 18190 $n $u)" "$(req 18191 $n $u)" "$m"
done; done
