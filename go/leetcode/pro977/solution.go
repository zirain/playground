package pro977

func sortedSquares(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1
	ans := make([]int, len(nums))
	for pos := n - 1; pos >= 0; pos-- {
		if v, w := nums[l]*nums[l], nums[r]*nums[r]; v <= w {
			r--
			ans[pos] = w
		} else {
			l++
			ans[pos] = v
		}
	}

	return ans
}
