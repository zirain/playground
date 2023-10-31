package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

const (
	podNamespaceEnv     = "POD_NAMESPACE"
	requestPortSplit    = ":"
	requertHostSPlit    = "."
	requestAuthorityKey = ":AUTHORITY"
)

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
	// Embed the default VM context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultVMContext
}

// Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
	// the remaining token for rate limiting, refreshed periodically.
	remainToken int
	// // the preconfigured request per second for rate limiting.
	// requestPerSecond int
	// NOTE(jianfeih): any concerns about the threading and mutex usage for tinygo wasm?
	// the last time the token is refilled with `requestPerSecond`.
	lastRefillNanoSec int64
}

// Override types.DefaultPluginContext.
func (p *pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpContext{contextID: contextID, pluginContext: p}
}

type httpContext struct {
	// Embed the default http context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultHttpContext
	contextID     uint32
	pluginContext *pluginContext
}

var additionalHeaders = map[string]string{
	"x-envoy-wasm-plugin": "autocomplete-authority",
}

func (ctx *httpContext) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	for key, value := range additionalHeaders {
		proxywasm.AddHttpResponseHeader(key, value)
	}
	return types.ActionContinue
}

func (ctx *httpContext) OnHttpRequestHeaders(int, bool) types.Action {
	authority, _ := proxywasm.GetHttpRequestHeader(requestAuthorityKey)
	if authority == "" {
		proxywasm.LogErrorf("authority is empty")
		return types.ActionContinue
	}
	// TODO: WasmPlugin do not support set vm_config, we can not get env right now.
	ns := os.Getenv(podNamespaceEnv)
	if ns == "" {
		proxywasm.LogErrorf("pod namespace is empty")
		return types.ActionContinue
	}

	servers := strings.Split(authority, requestPortSplit)
	port := "80" // http default port
	if len(servers) > 1 {
		port = servers[1]
	}
	s := servers[0]
	serviceNames := strings.Split(s, requertHostSPlit)
	if len(serviceNames) == 1 {
		serviceNames = append(serviceNames, ns)
	}

	auth := fmt.Sprintf("%s:%s", strings.Join(serviceNames, requertHostSPlit), port)

	proxywasm.LogWarnf("authority change to %s", auth)
	proxywasm.ReplaceHttpRequestHeader(requestAuthorityKey, auth)

	return types.ActionContinue
}
