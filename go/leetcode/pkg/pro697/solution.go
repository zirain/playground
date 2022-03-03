package pro697

type entry struct {
	count, l, r int
}

func findShortestSubArray(nums []int) int {
	mp := make(map[int]entry, 0)
	for i, v := range nums {
		if e, has := mp[v]; has {
			e.count++
			e.r = i
			mp[v] = e
		} else {
			mp[v] = entry{1, i, i}
		}
	}

	maxCount := 0
	ans := 0
	for _, v := range mp {
		if v.count > maxCount {
			maxCount, ans = v.count, v.r-v.l+1
		} else if v.count == maxCount {
			ans = min(ans, v.r-v.l+1)
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
