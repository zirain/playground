package pro1748

import (
	"sort"
)

// 排序后，前后不想等的数即为唯一数
func sumOfUnique(nums []int) int {
	sort.Ints(nums)
	total := 0
	for i, v := range nums {
		if (i+1 < len(nums) && nums[i] == nums[i+1]) ||
			(i > 0 && nums[i] == nums[i-1]) {
			continue
		}

		total += v
	}

	return total
}
