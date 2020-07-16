package pro35

func searchInsert(nums []int, target int) int {
	total := len(nums)
	if total == 0 ||
		target <= nums[0] {
		return 0
	}

	if target > nums[total-1] {
		return total
	}

	l, r := 0, total-1
	for l < r && r-l > 1 {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid
		}
	}

	return l + 1
}
