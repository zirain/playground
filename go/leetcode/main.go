package main

import "fmt"

func rob(nums []int) int {
	currMax := 0
	prevMax := 0
	for _, val := range nums {
		tmp := currMax
		currMax = max(prevMax+val, currMax)
		prevMax = tmp
	}

	return currMax
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{2, 1, 4, 5, 3, 1, 1, 3}))
}
