cmake_minimum_required(VERSION 3.15)
STRING( REGEX REPLACE ".*/(.*)" "\\1" CURRENT_FOLDER ${CMAKE_CURRENT_SOURCE_DIR} )
project(istio-proxy)

macro(print_all_variables)
    message(STATUS "print_all_variables------------------------------------------{")
    get_cmake_property(_variableNames VARIABLES)
    foreach (_variableName ${_variableNames})
        message(STATUS "${_variableName}=${${_variableName}}")
    endforeach()
    message(STATUS "print_all_variables------------------------------------------}")
endmacro()

set(CMAKE_CXX_STANDARD 17)
add_definitions(-DNULL_PLUGIN) # enable wasm nullvm navigation

file(STRINGS sourcefiles.txt all_SRCS)

message(STATUS "CMAKE_SOURCE_DIR=${CMAKE_SOURCE_DIR}")
message(STATUS "CMAKE_HOME_DIRECTORY=${CMAKE_HOME_DIRECTORY}")

add_executable(istio-proxy ${all_SRCS})

set(istio_include_dirs
        "./"
        "./src"
        "./extensions"

        "./bazel-${CURRENT_FOLDER}/external/envoy"
        "./bazel-${CURRENT_FOLDER}/external/envoy/source"
        "./bazel-${CURRENT_FOLDER}/external/envoy/include"
        "./bazel-${CURRENT_FOLDER}/external/envoy/api/wasm/cpp"
        "./bazel-${CURRENT_FOLDER}/external/boringssl/src/include/"
        "./bazel-${CURRENT_FOLDER}/external/com_github_gabime_spdlog/include"
        "./bazel-${CURRENT_FOLDER}/external/com_github_c_ares_c_ares"
        "./bazel-${CURRENT_FOLDER}/external/com_google_absl"
        "./bazel-${CURRENT_FOLDER}/external/com_google_cel_cpp"
        "./bazel-${CURRENT_FOLDER}/external/com_google_protobuf/src"
        "./bazel-${CURRENT_FOLDER}/external/com_github_fmtlib_fmt/include"
        "./bazel-${CURRENT_FOLDER}/external/com_github_eile_tclap/include"
        "./bazel-${CURRENT_FOLDER}/external/com_github_grpc_grpc/include"
        "./bazel-${CURRENT_FOLDER}/external/com_envoyproxy_protoc_gen_validate/"
        "./bazel-${CURRENT_FOLDER}/external/com_github_tencent_rapidjson/include/"
        "./bazel-${CURRENT_FOLDER}/external/com_github_datadog_dd_opentracing_cpp/include/"
        "./bazel-${CURRENT_FOLDER}/external/com_github_libevent_libevent/include"
        "./bazel-${CURRENT_FOLDER}/external/com_github_mirror_tclap/include"
        "./bazel-${CURRENT_FOLDER}/external/com_github_grpc_grpc"
        "./bazel-${CURRENT_FOLDER}/external/com_github_circonus_labs_libcircllhist/src/"
        "./bazel-${CURRENT_FOLDER}/external/com_github_nodejs_http_parser"
        "./bazel-${CURRENT_FOLDER}/external/com_github_nghttp2_nghttp2/lib/includes/"
        "./bazel-${CURRENT_FOLDER}/external/com_github_cyan4973_xxhash/"
        "./bazel-${CURRENT_FOLDER}/external/com_github_google_flatbuffers/include/"
        "./bazel-${CURRENT_FOLDER}/external/com_github_fmtlib_fmt/test"

        "./bazel-bin"
        "./bazel-bin/external/envoy_api"
        "./bazel-bin/external/mixerapi_git"
        "./bazel-bin/external/com_envoyproxy_protoc_gen_validate"
        "./bazel-bin/external/com_google_googleapis"
        "./bazel-bin/external/com_github_cncf_udpa"
)

target_include_directories(istio-proxy PUBLIC ${istio_include_dirs})