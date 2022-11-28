package main

import (
	"fmt"

	workloadv1alpha1 "github.com/karmada-io/karmada/examples/customresourceinterpreter/apis/workload/v1alpha1"
	configv1alpha1 "github.com/karmada-io/karmada/pkg/apis/config/v1alpha1"
	"github.com/karmada-io/karmada/pkg/resourceinterpreter/customizedinterpreter/webhook"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func main() {

}

func Interpreter(req *webhook.RequestAttributes) *webhook.ResponseAttributes {
	switch req.Operation {
	case configv1alpha1.InterpreterOperationInterpretReplica:
		return responseWithExploreReplica(req)
	default:
		return nil
	}
}

func responseWithExploreReplica(req *webhook.RequestAttributes) *webhook.ResponseAttributes {
	wantedWorkload := &workloadv1alpha1.Workload{}
	if err := convertToTypedObject(req.Object, wantedWorkload); err != nil {
		return nil
	}
	return &webhook.ResponseAttributes{
		Replicas: *wantedWorkload.Spec.Replicas + 1,
	}
}

// convertToTypedObject converts an unstructured object to typed.
func convertToTypedObject(in, out interface{}) error {
	if in == nil || out == nil {
		return fmt.Errorf("convert objects should not be nil")
	}

	switch v := in.(type) {
	case *unstructured.Unstructured:
		return runtime.DefaultUnstructuredConverter.FromUnstructured(v.UnstructuredContent(), out)
	case map[string]interface{}:
		return runtime.DefaultUnstructuredConverter.FromUnstructured(v, out)
	default:
		return fmt.Errorf("convert object must be pointer of unstructured or map[string]interface{}")
	}
}
