RM := rm -rf
WASM_PLUGIN_NAME := wasm-log.wasm

.PHONY : clean
clean:
	$(RM) ./$(WASM_PLUGIN_NAME)

build: clean
	tinygo build -o $(WASM_PLUGIN_NAME) -scheduler=none -target=wasi main.go

run:
	envoy -c envoy.yaml --component-log-level filter:debug,router:debug;wasm:debug