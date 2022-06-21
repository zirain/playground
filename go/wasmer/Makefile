
.PHONY: build-wasm-tinygo
build-wasm-tinygo:
	tinygo build -o sum.wasm -scheduler=none -target=wasi sum/main.go

.PHONY: run-go
run-tinygo: build-wasm-tinygo
	go run simple.go

.PHONY: run-node
run-node: build-wasm-tinygo
	node --experimental-wasi-unstable-preview1 simple.js
