package main

import (
	"fmt"
	"sort"
)

func minIncrementForUnique(A []int) int {

	sort.Ints(A)
	ans, taken := 0, 0

	for i := 1; i < len(A); i++ {
		if A[i-1] == A[i] {
			taken++
			ans -= A[i]
		} else {
			var give = min(taken, A[i]-A[i-1]-1)
			ans += give*(give+1)/2 + give*A[i-1]
			taken -= give
		}
	}
	if len(A) > 0 {
		ans += taken*(taken+1)/2 + taken*A[len(A)-1]
	}

	return ans
}

func min(x int, y int) int {
	if x < y {
		return x
	}

	return y
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}
