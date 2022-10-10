#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

LOGGING_LEVEL=debug kurator install istio --primary member1 --remote member2 --set meshConfig.accessLogFile=/dev/stdout