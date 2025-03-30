FROM ubuntu:bionic AS clang_context_amd64
ENV UBUNTU_RELEASE_CODE_NAME=bionic
FROM ubuntu:focal AS clang_context_arm64
ENV UBUNTU_RELEASE_CODE_NAME=focal
# hadolint ignore=DL3006
FROM clang_context_${TARGETARCH} AS clang_context

# hadolint ignore=DL3008
RUN apt-get update && apt-get install -y --no-install-recommends \
    xz-utils \
    wget \
    ca-certificates \
    software-properties-common \
    curl \
    ninja-build \
    gpg-agent \
    libtinfo5

ENV LLVM_VERSION=18.1.8
ENV LLVM_BASE_URL=https://github.com/llvm/llvm-project/releases/download/llvmorg-${LLVM_VERSION}
ENV LLVM_DIRECTORY=/usr/lib/llvm

RUN set -eux; \
    \
    case $(uname -m) in \
        x86_64) \
               LLVM_ARCHIVE=clang+llvm-${LLVM_VERSION}-x86_64-linux-gnu-ubuntu-18.04 \
               LLVM_ARTIFACT=clang+llvm-${LLVM_VERSION}-x86_64-linux-gnu-ubuntu-18.04;; \
        aarch64)  \
               LLVM_ARCHIVE=clang+llvm-${LLVM_VERSION}-aarch64-linux-gnu \
               LLVM_ARTIFACT=clang+llvm-${LLVM_VERSION}-aarch64-linux-gnu;; \
        *) echo "unsupported architecture"; exit 1 ;; \
    esac; \
    \
    wget -nv ${LLVM_BASE_URL}/${LLVM_ARTIFACT}.tar.xz; \
    tar -xJf ${LLVM_ARTIFACT}.tar.xz -C /tmp; \
    mkdir -p ${LLVM_DIRECTORY}; \
    mv /tmp/${LLVM_ARCHIVE}/* ${LLVM_DIRECTORY}/

# CMake >=3.20.0
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
# ubuntu 18.04 without arm64 version see https://apt.kitware.com/
# hadolint ignore=DL4001
RUN curl -fsSL https://apt.kitware.com/keys/kitware-archive-latest.asc | apt-key add -
RUN apt-add-repository "deb https://apt.kitware.com/ubuntu/ ${UBUNTU_RELEASE_CODE_NAME} main"
RUN apt-get update && apt-get install -y --no-install-recommends cmake

# RUN wget -qO- https://github.com/Kitware/CMake/releases/download/v3.31.6/cmake-3.31.6-linux-$(uname -m).tar.gz | tar --strip-components=1 -xz -C /usr/local

# TSan instrumented libc++.
ENV PATH=${LLVM_DIRECTORY}/bin:$PATH

WORKDIR /tmp
COPY proxy-tsan-instrumented-libcxx.sh proxy-tsan-instrumented-libcxx.sh
RUN ./proxy-tsan-instrumented-libcxx.sh