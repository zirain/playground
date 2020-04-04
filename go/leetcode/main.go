package main

import "fmt"

func trap(height []int) int {
	ans := 0
	size := len(height)
	for i := 1; i < size-1; i++ {
		max_left, max_right := 0, 0
		for j := i; j >= 0; j-- { //Search the left part for max bar size
			max_left = max(max_left, height[j])
		}
		for j := i; j < size; j++ { //Search the right part for max bar size
			max_right = max(max_right, height[j])
		}
		ans += min(max_left, max_right) - height[i]
	}
	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
