package main

import (
	"encoding/json"
	"fmt"
)

type demo struct {
	ClusterId  string `json:"clusterId"`
	ClusterIdX string `json:"cluster_id"`
}

func MarshalDemo() {
	d := demo{
		ClusterId:  "xxxx",
		ClusterIdX: "xxxx",
	}

	b, _ := json.Marshal(d)

	fmt.Println(string(b))
}

func main() {
	MarshalDemo()
}
