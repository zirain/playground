package pro35

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		n := nums[mid]

		if n == target {
			return mid
		} else if n > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
