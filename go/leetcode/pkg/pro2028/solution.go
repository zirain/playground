package pro2028

func missingRolls(rolls []int, mean int, n int) []int {

	m := len(rolls)
	total := mean * (m + n)

	mTotal := 0
	for _, r := range rolls {
		mTotal += r
	}
	left, i := total-mTotal, 0
	if left <= 0 {
		return []int{}
	}

	avg := left / n
	if avg > 6 || avg == 0 {
		return []int{}
	}

	ans := make([]int, n)
	for i := range ans {
		left -= avg
		ans[i] = avg
	}

	for left > 0 {
		ans[i]++
		if ans[i] > 6 {
			return []int{}
		}

		left--
		i++
		if i >= len(ans) {
			i = 0
		}
	}

	if left > 0 {
		return []int{}
	}

	return ans
}
