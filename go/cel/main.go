package main

import (
	"fmt"

	"github.com/google/cel-go/cel"
)

func main() {
	env, _ := cel.NewEnv()

	celTexts := []string{
		"response.code >= 500", // valid
		"code >= 500",          // valid
		")++++",                // invalid
	}

	for _, txt := range celTexts {
		_, issue := env.Parse(txt)
		if issue.Err() != nil {
			fmt.Printf("cel parse failed. txt: %s err: %v \n", txt, issue.Err())
			return
		}
	}

	fmt.Println("all cel parse success.")
}
