#!/bin/bash

set -x

PROTOC_VERSION=21.9
PROTOC_ZIP=protoc-${PROTOC_VERSION}-linux-x86_64.zip;
cd /tmp
wget -nv -O "/tmp/${PROTOC_ZIP}" "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}"
unzip "/tmp/${PROTOC_ZIP}"
mv /tmp/bin/protoc /usr/local/bin/
chmod +x /usr/local/bin/protoc