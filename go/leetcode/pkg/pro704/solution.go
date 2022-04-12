package pro704

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if l >= len(nums) || nums[l] != target {
		return -1
	}

	return l
}
