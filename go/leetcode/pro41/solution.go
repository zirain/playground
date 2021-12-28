package pro41

func firstMissingPositive(nums []int) int {
	dict := make(map[int]bool, 0)
	for _, val := range nums {
		dict[val] = true
	}

	ans := 1
	for {
		if _, ok := dict[ans]; ok {
			ans++
		} else {
			break
		}
	}

	return ans
}
