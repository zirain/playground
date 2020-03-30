package main

import (
	"fmt"
)

func lastRemaining(n int, m int) int {
	return f(n, m)
}

func f(n int, m int) int {
	if n == 1 {
		return 0
	}

	x := f(n-1, m)
	return (m + x) % n
}

func main() {
	fmt.Println(lastRemaining(5, 3))
	fmt.Println(lastRemaining(10, 17))
}
