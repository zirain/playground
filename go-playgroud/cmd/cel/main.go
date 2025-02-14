package main

import (
	"fmt"

	"github.com/google/cel-go/cel"
	"sigs.k8s.io/yaml"
)

func main() {
	celDemo1()

	expr := "string(node.metadata.INTERCEPTION_MODE)"
	out, err := getNodeMetadata(expr)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(out)
}

func celDemo1() {
	env, _ := cel.NewEnv()

	celTexts := []string{
		"xds.node.id == \"node_name\"",
		"response.code >= 500", // valid
		"code >= 500",          // valid
		"code == 200",          // valid
		")++++",                // invalid
	}

	for _, txt := range celTexts {
		ast, issue := env.Parse(txt)
		if issue.Err() != nil {
			fmt.Printf("cel parse failed. txt: %s err: %v \n", txt, issue.Err())
			return
		} else {
			expr, _ := cel.AstToParsedExpr(ast)
			out, _ := yaml.Marshal(expr.GetExpr())
			fmt.Println(string(out))
		}
	}

	fmt.Println("all cel parse success.")
}

func getNodeMetadata(celExpr string) (interface{}, error) {
	//node := readNode()

	celEnv, err := cel.NewEnv()
	if err != nil {
		return nil, err
	}
	ast, iss := celEnv.Parse(celExpr)
	if iss.Err() != nil {
		return nil, iss.Err()
	}
	prg, err := celEnv.Program(ast)
	if err != nil {
		return nil, err
	}

	out, _, err := prg.Eval(map[string]interface{}{
		"node.metadata": map[string]string{
			"INTERCEPTION_MODE": "REDIRECT",
		},
	})
	if err != nil {
		return nil, err
	}
	return out.Value(), nil
}
