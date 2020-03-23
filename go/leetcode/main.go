package main

import "fmt"

func massage(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp0 := 0
	dp1 := nums[0]

	for i := 1; i < n; i++ {
		tdp0 := max(dp0, dp1)
		tdp1 := dp0 + nums[i]

		dp0 = tdp0
		dp1 = tdp1
	}

	return max(dp0, dp1)
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

func main() {
	fmt.Println(massage([]int{1, 2, 3, 1}))
	fmt.Println(massage([]int{2, 7, 9, 3, 1}))
	fmt.Println(massage([]int{2, 1, 4, 5, 3, 1, 1, 3}))
}
