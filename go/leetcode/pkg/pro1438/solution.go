package pro1438

func longestSubarray(nums []int, limit int) int {
	var minQ, maxQ []int
	ans := 0
	left := 0
	for right, v := range nums {
		for len(minQ) > 0 && minQ[len(minQ)-1] > v {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, v)
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < v {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, v)
		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0]-minQ[0] > limit {
			if nums[left] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[left] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
