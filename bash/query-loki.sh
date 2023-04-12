#!/bin/bash

set -e
set -u
set -x
set -o pipefail


function count_by_pod() {
  local namespace="$1"
  local name="$2"
  curl -G -s 'http://localhost:3100/loki/api/v1/query_range' --data-urlencode "query={namespace=\"$namespace\", pod=\"$name\"}" | jq '.data.result[0].values | length'
}

SLEEP_POD=$(kubectl get pod -l app=sleep -o jsonpath={.items..metadata.name})
HTTPBIN_POD=$(kubectl get pod -l app=httpbin -o jsonpath={.items..metadata.name})

C1=$(count_by_pod "default" "$SLEEP_POD")
count_by_pod "default" "$HTTPBIN_POD"

verify_same() {
  local actual="$1"
  local expected="$2"
  if  [[ "$actual" != "$expected" ]]; then
    exit 1
  fi
}
verify_same 1 0
verify_same 1 1
