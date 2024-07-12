package main

import (
	"os"

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
	"x-envoy-wasm-plugin": "display-metadata",
}

func (ctx *pluginContext) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
	_, err := proxywasm.GetPluginConfiguration()
	if err != nil && err != types.ErrorStatusNotFound {
		proxywasm.LogCriticalf("error reading plugin configuration: %v", err)
		return types.OnPluginStartStatusFailed
	}

	return types.OnPluginStartStatusOK
}

func (ctx *httpContext) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	proxywasm.LogErrorf("OnHttpResponseHeaders")
	for key, value := range additionalHeaders {
		proxywasm.AddHttpResponseHeader(key, value)
	}

	return types.ActionContinue
}

func (ctx *httpContext) OnHttpStreamDone() {
	proxywasm.LogErrorf("OnHttpStreamDone")

	reqHeaders, err := proxywasm.GetHttpRequestHeaders()
	if err != nil {
		proxywasm.LogErrorf("failed to get request headers: %v", err)
		return
	}
	for _, h := range reqHeaders {
		proxywasm.LogErrorf("request header: <%s: %s>", h[0], h[1])
	}
	return
}

func (ctx *httpContext) OnHttpRequestHeaders(int, bool) types.Action {
	proxywasm.LogErrorf("OnHttpRequestHeaders")

	ns := os.Getenv(podNamespaceEnv)
	if ns == "" {
		proxywasm.LogErrorf("pod namespace is empty")
		return types.ActionContinue
	}
	proxywasm.LogErrorf("pod namespace is %s", ns)

	return types.ActionContinue
}
