#!/bin/bash

set -euo pipefail

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
NAMESPACE=${NAMESPACE:-sample}
REVISION=${REVISION:-1-31-1}
TARGET=${TARGET:-http://172.18.255.201/get}
RATE=${RATE:-10}
DURATION=${DURATION:-30s}
RESULTS=${RESULTS:-results.bin}

# Send traffic through the gateway for the whole upgrade
echo "GET ${TARGET}" | vegeta attack -rate="${RATE}" -duration="${DURATION}" > "${RESULTS}" &
VEGETA_PID=$!
trap 'kill "${VEGETA_PID}" 2>/dev/null || true' EXIT

# Switch the ingress gateway to the canary revision
sed "s|^\([[:space:]]*istio.io/rev:\).*|\1 ${REVISION}|" "${SCRIPT_DIR}/gateway.yaml" | kubectl apply -f -
kubectl rollout status deployment/istio-ingressgateway -n istio-ingress --timeout=5m

# Point the sample namespace at the canary revision
# (istio-injection takes precedence over istio.io/rev, so drop it)
kubectl label namespace "${NAMESPACE}" istio-injection- istio.io/rev="${REVISION}" --overwrite

# Restart workloads so sidecars get re-injected with the new revision
kubectl rollout restart deployment -n "${NAMESPACE}"
for deploy in $(kubectl get deployment -n "${NAMESPACE}" -o name); do
  kubectl rollout status "${deploy}" -n "${NAMESPACE}" --timeout=5m
done

istioctl proxy-status -r "${REVISION}"

# Wait for the attack to finish, then report
wait "${VEGETA_PID}"
vegeta report < "${RESULTS}"
