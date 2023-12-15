# how to dev

## dev with vscode

install extesions:
- bazel-statck-vscode-cc
- clangd

```
bazel build --override_repository=bazel_vscode_compdb=/root/.vscode-server/extensions/stackbuild.bazel-stack-vscode-cc-1.2.0/compdb/ --aspects=@bazel_vscode_compdb//:aspects.bzl%compilation_database_aspect --color=no --show_result=2147483647 --noshow_progress --noshow_loading_progress --output_groups=compdb_files,header_files --build_event_json_file=/tmp/tmp-envoy envoy

/root/.vscode-server/extensions/stackbuild.bazel-stack-vscode-cc-1.2.0/compdb/postprocess.py -b /tmp/tmp-envoy && rm /tmp/tmp-envoy
```

## setup

https://cloud.tencent.com/developer/article/1996907

## build

```console
bazel/setup_clang.sh /usr/lib/llvm-14/
bazel build -c opt envoy
cp bazel-bin/source/exe/envoy-static /usr/local/bin/envoy-dev
```

## build with docker

```console
./ci/run_envoy_docker.sh './ci/do_ci.sh bazel.dev'
```

```
cd /root/go/src/istio.io/envoy
#bazel/setup_clang.sh /usr/lib/llvm-14/
sha256sum /usr/local/bin/envoy-dev
bazel build -c opt envoy && cp bazel-bin/source/exe/envoy-static /usr/local/bin/envoy-dev
sha256sum /usr/local/bin/envoy-dev
```

build istio proxy with local envoy:

```
cd ~/go/src/istio.io/proxy
BAZEL_BUILD_ARGS='--override_repository=envoy=/root/go/src/istio.io/envoy' make build 
cp bazel-bin/src/envoy/envoy /usr/local/bin/envoy-dev
```

## release

```
ci/run_envoy_docker.sh 'ci/do_ci.sh bazel.release'
```

## test

```console
bazel test //test/extensions/access_loggers/open_telemetry:grpc_access_log_impl_test

bazel test //test/common/router:router_ratelimit_test

bazel test //test/common/common:matchers_test

bazel test //test/extensions/tracers/datadog:datadog_tracer_impl_test

bazel test //test/common/stats:tag_extractor_impl_test
```

## coverage test

```console
# find result in /tmp/envoy-docker-build/envoy/x64/generated/coverage.tar.zst
rm -rf /tmp/envoy-docker-build/envoy/x64/testlogs.tar.zst /tmp/envoy-docker-build/envoy/x64/generated/coverage.tar.zst
./ci/run_envoy_docker.sh './ci/do_ci.sh bazel.coverage //test/extensions/formatter/cel:cel_test'
```

## make doc

***please commit your changes before run following command***

```console
ci/run_envoy_docker.sh 'ci/do_ci.sh docs'

# witout docker
docs/build.sh
```

## lint

```shell
tools/local_fix_format.sh $(git diff --name-only | grep -E '\.(h|c|cc|proto)$')

tools/local_fix_format.sh -all
```
