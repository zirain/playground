package pro34

func searchRange(nums []int, target int) []int {
	l, r := leftBound(nums, target), rightBound(nums, target)
	return []int{l, r}
}

func leftBound(nums []int, target int) int {
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

func rightBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if r < 0 || nums[r] != target {
		return -1
	}

	return r
}
