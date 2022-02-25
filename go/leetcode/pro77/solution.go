package pro77

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
func combine(n int, k int) [][]int {
	ans := combineBy(1, n, []int{}, k)
	return ans
}

func combineBy(start, end int, prefix []int, k int) [][]int {
	if k == 0 {
		return [][]int{prefix}
	}

	if end-start+1 < k {
		return [][]int{}
	}

	l := len(prefix)
	ans := make([][]int, 0)
	for i := start; i <= end; i++ {
		tmp := make([]int, l+1)
		copy(tmp, prefix)
		tmp[l] = i

		ans = merge(ans, combineBy(i+1, end, tmp, k-1))
	}

	return ans
}

func merge(grid1, grid2 [][]int) [][]int {
	for _, arr := range grid2 {
		grid1 = append(grid1, arr)
	}

	return grid1
}
