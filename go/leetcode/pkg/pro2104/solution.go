package pro2104

func subArrayRanges(nums []int) (ans int64) {
	for i, num := range nums {
		minVal, maxVal := num, num
		for _, v := range nums[i+1:] {
			minVal = min(minVal, v)
			maxVal = max(maxVal, v)
			ans += int64(maxVal - minVal)
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
