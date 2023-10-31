package main

import (
	"fmt"
	"go/ast"
	"sort"
	"unicode"

	"golang.org/x/tools/go/packages"
	meshconfig "istio.io/api/mesh/v1alpha1"
)

var _ = &meshconfig.MeshConfig{}

func filterServiceTypes() []string {
	ans := make([]string, 0)
	config := &packages.Config{
		Mode: packages.NeedSyntax,
	}
	pkgs, _ := packages.Load(config, "istio.io/api/mesh/v1alpha1")
	pkg := pkgs[0]

	for _, s := range pkg.Syntax {
		for n, o := range s.Scope.Objects {
			if o.Kind == ast.Typ {
				// check if type is exported(only need for non-local types)
				if unicode.IsUpper([]rune(n)[0]) {
					node := o.Decl
					switch node.(type) {
					case *ast.TypeSpec:
						typeSpec := node.(*ast.TypeSpec)
						switch typeSpec.Type.(type) {
						case *ast.StructType:
							structType := typeSpec.Type.(*ast.StructType)
							for _, field := range structType.Fields.List {
								if i, ok := field.Type.(*ast.Ident); ok {
									fieldType := i.Name
									if fieldType != "string" {
										continue
									}

									for _, name := range field.Names {
										if name.Name == "Service" {
											ans = append(ans, n)
										}
									}
								}
							}
						}

					}

				}
			}
		}
	}

	return ans
}

func main() {
	types := filterServiceTypes()
	sort.Strings(types)

	fmt.Println("the following types contains `Service` field:")
	for _, t := range types {
		fmt.Println(t)
	}

}
