package main

import (
	"errors"
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
	if err != nil && !errors.Is(err, types.ErrorStatusNotFound) {
		proxywasm.LogCriticalf("error reading plugin configuration: %v", err)
		return types.OnPluginStartStatusFailed
	}

	return types.OnPluginStartStatusOK
}

func (ctx *httpContext) OnHttpResponseBody(_ int, _ bool) types.Action {
	responseCode, err := proxywasm.GetProperty([]string{"response", "code"})
	if err != nil {
		proxywasm.LogCriticalf("error getting response code: %v", err)
		return types.ActionContinue
	}

	switch string(responseCode) {
	case "403":
		if err := proxywasm.ReplaceHttpResponseBody([]byte("")); err != nil {
			proxywasm.LogCriticalf("error replacing response body: %v", err)
			return types.ActionContinue
		}
	}

	return types.ActionContinue
}
