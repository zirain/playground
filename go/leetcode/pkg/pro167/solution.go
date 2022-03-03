package pro167

func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		t := nums[l] + nums[r]
		if t == target {
			return []int{l + 1, r + 1}
		} else if t < target {
			l++
		} else {
			r--
		}

	}

	return []int{l, r}
}
