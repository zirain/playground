#!/bin/bash

set -x


LLVM_VERSION=12.0.1
LLVM_BASE_URL=https://github.com/llvm/llvm-project/releases/download/llvmorg-${LLVM_VERSION}
LLVM_DIRECTORY=/usr/lib/llvm
LLVM_ARCHIVE=clang+llvm-${LLVM_VERSION}-x86_64-linux-gnu-ubuntu-
LLVM_ARTIFACT=clang+llvm-${LLVM_VERSION}-x86_64-linux-gnu-ubuntu-16.04


wget -nv ${LLVM_BASE_URL}/${LLVM_ARTIFACT}.tar.xz
tar -xJf ${LLVM_ARTIFACT}.tar.xz -C /tmp
mkdir -p ${LLVM_DIRECTORY}
mv /tmp/${LLVM_ARCHIVE}/* ${LLVM_DIRECTORY}
rm -rf ${LLVM_BASE_URL}/${LLVM_ARTIFACT}.tar.xz

echo "${LLVM_DIRECTORY}/lib" | tee /etc/ld.so.conf.d/llvm.conf
ldconfig