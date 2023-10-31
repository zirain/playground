package main

func main() {}

// WARNING: wasmer does not support multi-returns
//export Sum
func Sum(x, y int) int {
	return x + y
}
