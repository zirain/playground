package main

import (
	"fmt"
	"log"
	"os"

	configv1alpha1 "github.com/karmada-io/karmada/pkg/apis/config/v1alpha1"
	wasmer "github.com/wasmerio/wasmer-go/wasmer"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type RequestAttributes struct {
	Operation   configv1alpha1.InterpreterOperation
	Object      *unstructured.Unstructured
	ObservedObj *unstructured.Unstructured
	ReplicasSet int32
}

func main() {
	attributes := &RequestAttributes{}

	wasmBytes, err := os.ReadFile("/etc/karmada/interpreter.wasm")
	if err != nil {
		log.Fatalf("load wasmer module err: %v", err)
		return
	}

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
	importObject, err := wasiEnv.GenerateImportObject(store, module)
	if err != nil {
		log.Fatalf("generate wasmer objecer err: %v", err)
		return
	}
	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		log.Fatalf("new wasmer instance err: %v", err)
		return
	}

	// Gets the `sum` exported function from the WebAssembly instance.
	fn, err := instance.Exports.GetFunction("Interpreter")
	if err != nil {
		log.Fatalf("get wasmer module err: %v", err)
		return
	}

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, _ := fn(attributes)
	// attrs, ok := result.(*webhook.ResponseAttributes)
	// if !ok {
	// 	return
	// }

	fmt.Println(result) // 42!
}
