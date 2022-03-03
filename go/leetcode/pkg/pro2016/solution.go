package pro2016

func maximumDifference(nums []int) int {
	var (
		n   = len(nums)
		ans = -1
	)

	for l := 0; l <= n-1; l++ {
		for r := l + 1; r < n; r++ {
			t := nums[r] - nums[l]
			if t > 0 {
				ans = max(ans, t)
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
