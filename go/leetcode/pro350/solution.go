package pro350

func intersect(nums1 []int, nums2 []int) []int {
	dict := make(map[int]int, 0)
	for _, val := range nums1 {
		dict[val]++
	}

	ans := make([]int, 0)
	for _, val := range nums2 {
		if _, ok := dict[val]; !ok {
			continue
		}

		if dict[val] > 0 {
			ans = append(ans, val)
			dict[val]--
		}
	}
	return ans
}
