package pro1

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, n := range nums {
		left := target - n
		if idx, ok := dict[left]; ok {
			return []int{i, idx}
		} else {
			dict[n] = i
		}
	}

	return []int{-1, -1}
}
