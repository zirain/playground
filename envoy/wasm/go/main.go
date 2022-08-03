package main

import (
	"encoding/binary"
	"time"

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

// Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
}

// Override types.DefaultPluginContext.
func (*pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpHeaders{contextID: contextID}
}

type httpHeaders struct {
	// Embed the default http context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultHttpContext
	contextID uint32
}

func (ctx *httpHeaders) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	path := []string{"request", "time"}
	startTime, err := proxywasm.GetProperty(path)

	if err != nil {
		proxywasm.LogCriticalf("GetProperty: %v", err)
	}

	unix_nano := binary.LittleEndian.Uint64(startTime) // 小字节序
	proxywasm.LogCriticalf("startTime UnixNano: %d", unix_nano)
	t := time.UnixMicro(int64(unix_nano) / 1000)
	proxywasm.LogCriticalf("startTime: %v", t)

	hs, err := proxywasm.GetHttpRequestHeaders()
	if err != nil {
		proxywasm.LogCriticalf("failed to get request headers: %v", err)
	}

	for _, h := range hs {
		proxywasm.LogInfof("request header --> %s: %s", h[0], h[1])
	}
	return types.ActionContinue
}

// Override types.DefaultHttpContext.
func (ctx *httpHeaders) OnHttpStreamDone() {
	path := []string{"request", "time"}
	startTime, err := proxywasm.GetProperty(path)

	if err != nil {
		proxywasm.LogCriticalf("GetProperty: %v", err)
	}

	proxywasm.LogCriticalf("startTime: %v", startTime)

	proxywasm.LogCriticalf("%d finished", ctx.contextID)
}
