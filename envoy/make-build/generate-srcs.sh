#!/bin/bash

set -ex

bazel_dir="bazel-${PWD##*/}"

find -L -E $bazel_dir/external src extensions -regex '.*\.(cc|c|cpp)' > sourcefiles.txt