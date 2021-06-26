package main

import (
	"encoding/json"
	"fmt"
)

type demo struct {
	ClusterId  string `json:"clusterId,omitempty"`
	ClusterIdX string `json:"cluster_id,omitempty"`
}

func MarshalDemo() {
	d := demo{
		ClusterId:  "xxxx",
		ClusterIdX: "xxxx",
	}

	b, _ := json.Marshal(d)
	fmt.Println(string(b))

	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("err: %w", err)
	}

	for k, v := range m {
		fmt.Printf("k: %s, v: %v\n", k, v)
	}
}

func main() {
	MarshalDemo()
}
