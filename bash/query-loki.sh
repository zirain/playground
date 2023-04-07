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

count_by_pod "default" "sleep-78ff5975c6-bj745"
count_by_pod "default" "helloworld-v1-78b9f5c87f-kn5vd"