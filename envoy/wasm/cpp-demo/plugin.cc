#include <string>
#include <unordered_map>

//#include <google/protobuf/util/time_util.h>
#include "proxy_wasm_intrinsics.h"

// Allows mapping from google::protobuf::util to other util libraries.
namespace ProtobufUtil = google::protobuf::util;

class ExampleContext : public Context {
public:
  explicit ExampleContext(uint32_t id, RootContext* root) : Context(id, root) {}

  FilterHeadersStatus onRequestHeaders(uint32_t headers, bool end_of_stream) override;
  void onDone() override;
  void onLog() override;
};
static RegisterContextFactory register_ExampleContext(CONTEXT_FACTORY(ExampleContext));

FilterHeadersStatus ExampleContext::onRequestHeaders(uint32_t headers, bool end_of_stream) {
  logInfo(std::string("onRequestHeaders ") + std::to_string(id()));
  auto path = getRequestHeader(":path");
  logInfo(std::string("header path ") + std::string(path->view()));
  return FilterHeadersStatus::Continue;
}

void ExampleContext::onDone() { logInfo("onDone " + std::to_string(id())); }

void ExampleContext::onLog() {
  // uint64_t unix_micros = 0;
  // if (getValue({"request", "time"}, &unix_micros)){
  //   auto request_time = ProtobufUtil::TimeUtil::NanosecondsToTimestamp(unix_micros);
  //   LOG_WARN(std::string("start_time @ " + ProtobufUtil::TimeUtil::ToString(request_time)));
  // }

  auto request_id = getProperty({"request", "id"});
  if (request_id.has_value()){
    LOG_WARN(std::string("request id @ " + request_id.value()->toString()));
  }

  auto request_user = getProperty({"request", "headers", "x-real-user"});
  if (request_user.has_value()){
    LOG_WARN(std::string("request user @ " + request_user.value()->toString()));
  }

  LOG_WARN(std::string("id " + std::to_string(id())));
}