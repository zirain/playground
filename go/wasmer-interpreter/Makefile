WASM_MODULE_NAME=interpreter.wasm
TINYGO=tinygo-dev

.PHONY: build-wasm
build-wasm:
	rm -rf $(WASM_MODULE_NAME)
	#GOOS=js GOARCH=wasm go build -o $(BIN_PATH)/$(WASM_MODULE_NAME) main.go
	CGO_ENABLED=0 $(TINYGO) build -o $(WASM_MODULE_NAME) -scheduler=none -tags=appengine -target=wasi interpreter/main.go
