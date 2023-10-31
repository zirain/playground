package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg := []string{
		"(response_code=\\.=(.+?);\\.;)|_rq(_(\\d{3}))$",
		"_rq(_(\\d{3}))$|(response_code=\\.=(.+?);\\.;)",
		"_rq(_(\\d{3}))$",
		"(response_code=\\.=(.+?);\\.;)",
	}

	buf := "cluster.xds-grpc.upstream_rq_200" // cluster.xds-grpc.upstream_rq{response_code="200"}
	for _, rString := range reg {
		r := regexp.MustCompile(rString)
		res := r.FindAllStringSubmatch(buf, -1)
		fmt.Println("regexp: ", rString, "res: ", len(res))
		for _, s := range res {
			fmt.Printf("len: %d, %v\n", len(s), s)
		}
	}
}
