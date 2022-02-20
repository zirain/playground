package pro3

func lengthOfLongestSubstring(s string) int {
	r, ans, n := -1, 0, len(s)

	m := make(map[byte]int)
	for l := 0; l < n; l++ {
		if l != 0 {
			delete(m, s[l-1])
		}

		for r+1 < n && m[s[r+1]] == 0 {
			m[s[r+1]]++
			// 不断地移动右指针
			r++
		}

		ans = max(ans, r-l+1)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
