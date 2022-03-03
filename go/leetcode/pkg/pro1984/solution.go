package pro1984

import "sort"

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	ans := nums[len(nums)-1] - nums[0]
	for i := 0; i <= len(nums)-k; i++ {
		diff := nums[i+k-1] - nums[i]
		ans = min(ans, diff)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
