#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail


docker rm -f primary-control-plane remote1-control-plane remote2-control-plane
docker volume prune -f
docker image prune -f
