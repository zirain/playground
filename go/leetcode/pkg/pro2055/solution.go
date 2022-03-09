package pro2055

func platesBetweenCandles(s string, queries [][]int) []int {
	ans := make([]int, len(queries))
	buf := []byte(s)

	for idx, rows := range queries {
		if rows[1]-rows[0] < 2 {
			continue
		}

		l, r := -1, -1
		for i := rows[0]; i < rows[1]; i++ {
			if buf[i] == '|' {
				l = i
				break
			}
		}

		for i := rows[1]; i > rows[0]; i-- {
			if buf[i] == '|' {
				r = i
				break
			}
		}

		if l >= 0 && r >= 0 && l < r {
			cnt := 0
			for i := l + 1; i < r; i++ {
				if buf[i] == '*' {
					cnt++
				}
			}

			ans[idx] = cnt
		}
	}

	return ans
}
