package main

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

var (
	errHexTraceIDWrongLen = errors.New("hex traceId span has wrong length (expected 16 or 32)")
	errHexTraceIDZero     = errors.New("traceId is zero")
)

func main() {
	traceID, err := hexToTraceID("404142434445464748494a4b4c4d4e4f")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(traceID[:]))

	traceID, err = hexToTraceID("00000000000000000ccce09bf12e94df")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(traceID[:]))
}

func hexToTraceID(hexStr string) (pcommon.TraceID, error) {
	// Per info at https://zipkin.io/zipkin-api/zipkin-api.yaml it should be 16 or 32 characters
	hexLen := len(hexStr)
	if hexLen != 16 && hexLen != 32 {
		return pcommon.NewTraceIDEmpty(), errHexTraceIDWrongLen
	}

	var id [16]byte
	if hexLen == 16 {
		if _, err := hex.Decode(id[8:], []byte(hexStr)); err != nil {
			return pcommon.NewTraceIDEmpty(), err
		}
		if pcommon.TraceID(id).IsEmpty() {
			return pcommon.NewTraceIDEmpty(), errHexTraceIDZero
		}
		return id, nil
	}

	if _, err := hex.Decode(id[:], []byte(hexStr)); err != nil {
		return pcommon.NewTraceIDEmpty(), err
	}
	if pcommon.TraceID(id).IsEmpty() {
		return pcommon.NewTraceIDEmpty(), errHexTraceIDZero
	}
	return id, nil
}
