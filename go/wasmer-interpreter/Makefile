WASM_MODULE_NAME=interpreter.wasm
TINYGO=tinygo-dev
TINYGO_DOCKER=ghcr.io/tinygo-org/tinygo-dev:sha-ee3af40cab35ba2b8241138e67d3dc2fb7da4406

.PHONY: build-wasm
build:
	rm -rf $(WASM_MODULE_NAME)
	#GOOS=js GOARCH=wasm go build -o $(BIN_PATH)/$(WASM_MODULE_NAME) main.go
	CGO_ENABLED=0 $(TINYGO) build -o $(WASM_MODULE_NAME) -scheduler=none -tags=appengine -target=wasi interpreter/main.go

.PHONY: docker-build
docker-build:
	docker run --rm -v $(pwd):/src $(TINYGO_DOCKER) tinygo build -o $(WASM_MODULE_NAME) -target=wasi interpreter/main.go
