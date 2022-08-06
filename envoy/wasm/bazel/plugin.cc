// Copyright Istio Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#include "absl/time/time.h"
#include "plugin.h"

FilterHeadersStatus ExampleContext::onRequestHeaders(uint32_t headers, bool end_of_stream) {
  logInfo(std::string("onRequestHeaders ") + std::to_string(id()));
  auto path = getRequestHeader(":path");
  logInfo(std::string("header path ") + std::string(path->view()));
  return FilterHeadersStatus::Continue;
}

void ExampleContext::onDone() { logInfo("onDone " + std::to_string(id())); }

void ExampleContext::onLog() {
  uint64_t timestamp;
  getValue({"request", "time"}, &timestamp);
  LOG_WARN(std::string("request time @ " + absl::FormatTime(absl::FromUnixNanos(timestamp))));


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