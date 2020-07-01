package main

import "fmt"

func minus(a, b int) int {
	return a - b
}

func add(a, b int) int {
	return a + b
}

type mathFunc func(int, int) int

func clac(a, b int, f mathFunc) int {
	return f(a, b)
}

func main() {
	var f mathFunc

	f = add
	fmt.Println(f(1, 2))

	f = minus
	fmt.Println(clac(1, 2, f))
}
