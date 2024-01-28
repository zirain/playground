package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net/http"

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
	customMessage map[string]string
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
	data, err := proxywasm.GetPluginConfiguration()
	if err != nil && err != types.ErrorStatusNotFound {
		proxywasm.LogCriticalf("error reading plugin configuration: %v", err)
		return types.OnPluginStartStatusFailed
	}

	var message map[string]string
	if err := json.Unmarshal(data, &message); err == nil {
		ctx.customMessage = message
	} else {
		proxywasm.LogCriticalf("error unmarshal configuration: %v", err)
	}

	return types.OnPluginStartStatusOK
}

func (ctx *httpContext) OnHttpResponseBody(bodySize int, endOfStream bool) types.Action {
	if !endOfStream {
		// Wait until we see the entire body to replace.
		return types.ActionPause
	}

	bs, err := proxywasm.GetProperty([]string{"response", "code"})
	if err != nil {
		proxywasm.LogCriticalf("error getting response code: %v", err)
		return types.ActionContinue
	}

	originalBody, err := proxywasm.GetHttpResponseBody(0, bodySize)
	if err != nil {
		proxywasm.LogCriticalf("failed to get response body: %v", err)
		return types.ActionContinue
	}

	code := binary.LittleEndian.Uint64(bs)
	codeStr := fmt.Sprintf("%d", code)
	switch code {
	case http.StatusForbidden:
		if string(originalBody) != "RBAC: access denied" {
			proxywasm.LogCriticalf("unexpected response body: %s", string(originalBody))
			return types.ActionContinue
		}

		msg := ctx.pluginContext.customMessage[codeStr]
		proxywasm.LogCriticalf("replace response body to: %s", msg)
		if err := proxywasm.ReplaceHttpResponseBody([]byte(msg + "\n")); err != nil {
			proxywasm.LogCriticalf("error replacing response body: %v", err)
			return types.ActionContinue
		}
	}

	return types.ActionContinue
}
