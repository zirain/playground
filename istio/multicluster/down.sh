#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail


kind delete clusters primary remote1 remote2
docker volume prune -f
docker image prune -f
