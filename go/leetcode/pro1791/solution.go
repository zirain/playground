package pro1791

// 实际不需要这么复杂，按照题意中间点的统计出现次数大于
func findCenter(edges [][]int) int {
	n := len(edges) + 1
	stats := make([]int, n)
	for _, edge := range edges {
		l, r := edge[0], edge[1]
		stats[l-1]++
		stats[r-1]++

		if stats[l-1] == n-1 {
			return l
		}

		if stats[r-1] == n-1 {
			return r
		}
	}

	// should not happen
	return 0
}
