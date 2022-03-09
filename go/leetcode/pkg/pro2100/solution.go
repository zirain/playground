package pro2100

func goodDaysToRobBank(security []int, time int) []int {
	ans := make([]int, 0)
	cnt := len(security)

	nonIncreaseDp := make([]int, cnt, cnt)
	for i, num := range security {
		if i != 0 && num <= security[i-1] {
			nonIncreaseDp[i] = nonIncreaseDp[i-1] + 1
		}
	}

	nonDecreaseDp := make([]int, cnt, cnt)
	for i := cnt - 1; i >= 0; i-- {
		if i != cnt-1 && security[i] <= security[i+1] {
			nonDecreaseDp[i] = nonDecreaseDp[i+1] + 1

		}
	}

	for i := 0; i < cnt; i++ {
		if nonIncreaseDp[i] >= time && nonDecreaseDp[i] >= time {
			ans = append(ans, i)
		}
	}
	return ans
}
