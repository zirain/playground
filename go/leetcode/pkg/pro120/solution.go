package pro120

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	prevs := make([]int, n)
	prevs[0] = triangle[0][0]

	for i := 1; i < n; i++ {
		k := len(triangle[i])
		cur := make([]int, n)
		for j := 0; j < k; j++ {
			if j == 0 {
				cur[j] = triangle[i][j] + prevs[j]
				continue
			}

			if j == k-1 {
				cur[j] = triangle[i][j] + prevs[j-1]
				continue
			}
			cur[j] = min(prevs[j-1]+triangle[i][j], prevs[j]+triangle[i][j])
		}
		copy(prevs, cur)
	}

	ans := prevs[0]
	for i := 1; i < len(prevs); i++ {
		ans = min(ans, prevs[i])
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
