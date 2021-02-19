package pro1004

import "sort"

func longestOnes(A []int, K int) int {
	ans := 0
	P := make([]int, len(A)+1)
	for i, v := range A {
		P[i+1] = P[i] + 1 - v
	}

	for right, v := range P {
		left := sort.SearchInts(P, v-K)
		ans = max(ans, right-left)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
