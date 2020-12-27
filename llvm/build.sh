#!/bin/bash

LLVM_VERSION=8.0.0
LLVM_SRC=/root/llvm-src
INSTALL_DIR=/usr/local/llvm-8

mkdir -p ${LLVM_SRC}/tool
mkdir -p ${LLVM_SRC}/llvm-src/tools/clang/tools


wget http://releases.llvm.org/${LLVM_VERSION}/llvm-${LLVM_VERSION}.src.tar.xz
wget http://releases.llvm.org/${LLVM_VERSION}/cfe-${LLVM_VERSION}.src.tar.xz
wget http://releases.llvm.org/${LLVM_VERSION}/clang-tools-extra-${LLVM_VERSION}.src.tar.xz
tar xvf llvm-${LLVM_VERSION}.src.tar.xz
tar xvf cfe-${LLVM_VERSION}.src.tar.xz
tar xvf clang-tools-extra-${LLVM_VERSION}.src.tar.xz
mv llvm-${LLVM_VERSION}.src ${LLVM_SRC}
mv cfe-${LLVM_VERSION}.src ${LLVM_SRC}/tool/clang
mv clang-tools-extra-${LLVM_VERSION}.src llvm-src/tools/clang/tools/extra


sudo mkdir -p ${INSTALL_DIR}
sudo mkdir -p ${LLVM_SRC}/build
cd ${LLVM_SRC}/build
cmake -G "Unix Makefiles" -DLLVM_TARGETS_TO_BUILD=X86 -DCMAKE_BUILD_TYPE="Release"  -DCMAKE_INSTALL_PREFIX="${INSTALL_DIR}" ..
make && make install