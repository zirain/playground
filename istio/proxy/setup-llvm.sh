#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail
set -x


LLVM_VERSION=12.0.0
LLVM_BASE_URL=https://github.com/llvm/llvm-project/releases/download/llvmorg-${LLVM_VERSION}
LLVM_DIRECTORY=/usr/lib/llvm-12
LLVM_ARCHIVE=clang+llvm-${LLVM_VERSION}-x86_64-linux-gnu-ubuntu-18.04
LLVM_ARTIFACT=clang+llvm-${LLVM_VERSION}-x86_64-linux-gnu-ubuntu-18.04

if [[ -f "${LLVM_ARTIFACT}.tar.xz" ]]; then
    echo "${LLVM_ARTIFACT}.tar.xz exists"
else
    wget -nv ${LLVM_BASE_URL}/${LLVM_ARTIFACT}.tar.xz
fi

tar -xJf ${LLVM_ARTIFACT}.tar.xz -C /tmp
rm -rf ${LLVM_DIRECTORY}
mkdir -p ${LLVM_DIRECTORY}
mv /tmp/${LLVM_ARCHIVE}/* ${LLVM_DIRECTORY}
rm -rf ${LLVM_ARTIFACT}.tar.xz

echo "${LLVM_DIRECTORY}/lib" | tee /etc/ld.so.conf.d/llvm.conf
ldconfig