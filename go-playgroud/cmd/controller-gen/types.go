// +k8s:deepcopy-gen=package
package main

type XdsResources map[string]int // working
//type XdsResources2 = map[string]int // not working: interface conversion: types.Type is *types.Map, not *types.Named
//type XdsResources3 map[string]interface{} // not working: invalid map value type: interface{}
