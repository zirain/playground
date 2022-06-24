# how to dev

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

## release

```
ci/run_envoy_docker.sh 'ci/do_ci.sh bazel.release'
```

## test

```console
bazel test //test/extensions/access_loggers/open_telemetry:grpc_access_log_impl_test

bazel test //test/common/router:router_ratelimit_test
```

## make doc

```console
ci/run_envoy_docker.sh 'ci/do_ci.sh docs'
```

## lint

```console
./ci/run_envoy_docker.sh './ci/do_ci.sh fix_format'

ci/run_envoy_docker.sh 'ci/do_ci.sh format_pre'

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