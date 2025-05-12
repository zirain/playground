
#!/usr/bin/env bash

set -euo pipefail

MIRROR_REGISTRY=${MIRROR_REGISTRY:-"192.168.4.146:5000"}
IMAGE_TAG=${IMAGE_TAG:-"latest"}
REPO_NAME=${REPO_NAME:-"envoyproxy"}
RMI_SOURCE=${RMI_SOURCE:-""}

exampleImages=(gateway-static-file-server
               gateway-preserve-case-backend
               gateway-grpc-ext-proc
               gateway-envoy-ext-auth
               gateway-extension-server
               gateway-simple-extension-server)
for imageName in ${exampleImages[@]} ; do
    SOURCE_IMAGE="${REPO_NAME}/${imageName}:${IMAGE_TAG}"
    MIRROR_IMAGE=${MIRROR_REGISTRY}/${REPO_NAME}/${imageName}:${IMAGE_TAG}
    docker tag "${SOURCE_IMAGE}" ${MIRROR_IMAGE}
    docker push ${MIRROR_IMAGE}
    docker rmi ${MIRROR_IMAGE}
    if [[ -n "${RMI_SOURCE}" ]]; then
        docker rmi ${SOURCE_IMAGE}
    fi
done