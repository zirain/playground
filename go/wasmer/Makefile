

.PHONY: build-wasm-sum
build-wasm-sum:
	tinygo build -o sum.wasm -scheduler=none -target=wasi sum/main.go

.PHONY: run-go
run-go: build-wasm-sum
	go run simple.go

.PHONY: run-node
run-node:
	node --experimental-wasi-unstable-preview1 simple.js
