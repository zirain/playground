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
bazel/setup_clang.sh /usr/lib/llvm-10/
bazel build -c opt envoy
cp bazel-bin/source/exe/envoy-static /usr/local/bin/envoy-dev
```

## build with docker

```console
./ci/run_envoy_docker.sh './ci/do_ci.sh bazel.dev'
```

```
cd /root/go/src/istio.io/envoy
#bazel/setup_clang.sh /usr/lib/llvm-10/
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
# check api
tools/proto_format/proto_format.sh

# check clang format
CLANG_FORMAT=clang-format ./tools/code_format/check_format.py fix
```

```console
./ci/run_envoy_docker.sh './ci/do_ci.sh fix_format'

ci/run_envoy_docker.sh 'ci/do_ci.sh format_pre'

./ci/run_envoy_docker.sh './ci/do_ci.sh fix_proto_format' && ./ci/run_envoy_docker.sh './ci/do_ci.sh format' && ./ci/run_envoy_docker.sh 'ci/check_and_fix_format.sh'
```

# envoy-handbook

https://jimmysong.io/envoy-handbook/


# 解析envoy处理http请求

https://github.com/dzdx/notes/tree/master/envoy


# 一个Http请求到响应的全链路
https://cloud.tencent.com/developer/article/1396880
https://cloud.tencent.com/developer/article/1396879
https://cloud.tencent.com/developer/article/1396877

# Envoy请求流程源码解析
- [流量劫持](https://zhuanlan.zhihu.com/p/471728761)
- [请求解析](https://zhuanlan.zhihu.com/p/475708734)