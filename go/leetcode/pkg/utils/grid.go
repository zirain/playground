package utils

import "fmt"

func RenderGrid(input [][]int) {
	fmt.Println("[")
	for i := 0; i < len(input); i++ {
		fmt.Println(" ", input[i])
	}
	fmt.Println("]")
}
