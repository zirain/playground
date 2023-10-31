package main

import (
	"fmt"
	"io/ioutil"
	"log"

	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	wasmBytes, _ := ioutil.ReadFile("sum.wasm")

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		log.Fatalf("new wasmer module err: %v", err)
		return
	}

	wasiEnv, _ := wasmer.NewWasiStateBuilder("wasi-program").Finalize()

	// Instantiates the module
	importObject, _ := wasiEnv.GenerateImportObject(store, module)
	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		log.Fatalf("new wasmer instance err: %v", err)
		return
	}

	// Gets the `sum` exported function from the WebAssembly instance.
	sum, err := instance.Exports.GetFunction("Sum")
	if err != nil {
		log.Fatalf("get wasmer module err: %v", err)
		return
	}

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, _ := sum(5, 37)

	fmt.Println(result) // 42!
}
