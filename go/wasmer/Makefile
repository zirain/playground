

.PHONY: build-wasm-sum
build-wasm-sum:
	tinygo build -o sum.wasm -scheduler=none -target=wasi sum/main.go

.PHONY: run
run: build-wasm-sum
	go run simple.go