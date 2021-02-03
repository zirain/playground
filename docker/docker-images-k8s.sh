#!/bin/bash

set -e
KUBE_VERSION=v1.19.3
KUBE_PAUSE_VERSION=3.2
ETCD_VERSION=3.4.13-0
COREDNS_VERSION=1.7.0
GCR_URL=k8s.gcr.io
ALIYUN_URL=registry.cn-hangzhou.aliyuncs.com/google_containers

# get images
images=(kube-proxy:${KUBE_VERSION}
    kube-scheduler:${KUBE_VERSION}
    kube-controller-manager:${KUBE_VERSION}
    kube-apiserver:${KUBE_VERSION}
    pause:${KUBE_PAUSE_VERSION}
    etcd:${ETCD_VERSION}
    coredns:${COREDNS_VERSION})

for imageName in ${images[@]} ; do
    docker pull $ALIYUN_URL/$imageName
    docker tag $ALIYUN_URL/$imageName $GCR_URL/$imageName
    docker rmi $ALIYUN_URL/$imageName
done

# show images
docker images