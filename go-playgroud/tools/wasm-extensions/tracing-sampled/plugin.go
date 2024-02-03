package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
	// Embed the default VM context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultVMContext
}

func (*vmContext) NewPluginContext(_ uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
}

func (ctx *pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpContext{contextID: contextID, pluginContext: ctx}
}

type httpContext struct {
	// Embed the default http context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultHttpContext
	contextID     uint32
	pluginContext *pluginContext
}

func (ctx *pluginContext) OnPluginStart(_ int) types.OnPluginStartStatus {
	_, err := proxywasm.GetPluginConfiguration()
	if err != nil && err != types.ErrorStatusNotFound {
		proxywasm.LogCriticalf("error reading plugin configuration: %v", err)
		return types.OnPluginStartStatusFailed
	}

	return types.OnPluginStartStatusOK
}

func (ctx *httpContext) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	reqPath, err := proxywasm.GetProperty([]string{"request", "path"})
	if err != nil {
		proxywasm.LogCriticalf("error reading request path: %v", err)
		return types.ActionContinue
	}

	if string(reqPath) == "/headers" {
		val, err := proxywasm.GetHttpRequestHeader("x-b3-sampled")
		if err != nil {
			proxywasm.LogCriticalf("error reading request header: %v", err)
		}

		proxywasm.LogCriticalf("x-b3-sampled: %s", val)
		proxywasm.ReplaceHttpRequestHeader("x-b3-sampled", "1")
		return types.ActionContinue
	}

	return types.ActionContinue
}

func (ctx *httpContext) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	//proxywasm.AddHttpResponseHeader("x-client-trace-id", "0")

	return types.ActionContinue
}
