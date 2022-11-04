#!/usr/bin/env bash

# Copyright Istio Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

WD=$(dirname "$0")
WD=$(cd "$WD"; pwd)

set -eux

# This script sets up the plain text rendered deployments for addons
# See samples/addons/README.md for more information

ADDONS="${WD}/addons"
DASHBOARDS="${WD}/dashboards"
mkdir -p "${ADDONS}"
LGTM_NS=${LGTM_NS:-"lgtm-stack"}
GRAFANA_VERSION=${GRAFANA_VERSION:-"6.31.1"}
LOKI_VERSION=${LOKI_VERSION:-"2.8.4"}
TEMPO_VERSION=${TEMPO_VERSION:-"0.16.3"}
TMP=$(mktemp -d)

function compressDashboard() {
  < "${DASHBOARDS}/$1" jq -c  > "${TMP}/$1"
}

# Set up grafana
{
  helm template grafana grafana \
    --namespace "${LGTM_NS}" \
    --version "${GRAFANA_VERSION}" \
    --repo https://grafana.github.io/helm-charts \
    -f "${WD}/values/grafana.yaml"
} > "${ADDONS}/grafana.yaml"

{
  helm template loki loki-stack \
    --namespace "${LGTM_NS}" \
    --version "${LOKI_VERSION}" \
    --repo https://grafana.github.io/helm-charts \
    -f "${WD}/values/loki.yaml"
} > "${ADDONS}/loki.yaml"

{
  helm template tempo tempo \
    --namespace "${LGTM_NS}" \
    --version "${TEMPO_VERSION}" \
    --repo https://grafana.github.io/helm-charts \
    -f "${WD}/values/tempo.yaml"
} > "${ADDONS}/tempo.yaml"