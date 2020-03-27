package main

import (
	"fmt"
)

func hasGroupsSizeX(deck []int) bool {
	dict := make(map[int]int)
	for _, val := range deck {
		if count, ok := dict[val]; ok {
			dict[val] = count + 1
		} else {
			dict[val] = 1
		}
	}

	g := -1
	for _, val := range dict {
		if g == -1 {
			g = val
		} else {
			g = gcd(val, g)
		}
	}

	return g >= 2
}

func gcd(x int, y int) int {
	if x == 0 {
		return y
	} else {
		return gcd(y%x, x)
	}
}

func main() {
	fmt.Println(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}))
	fmt.Println(hasGroupsSizeX([]int{1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1, 2, 2, 2, 2}))
}
