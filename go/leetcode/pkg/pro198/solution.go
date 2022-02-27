package pro198

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	pre, cur := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		t := max(pre+nums[i], cur)
		pre = cur
		cur = t
	}

	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
