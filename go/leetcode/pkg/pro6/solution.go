package pro6

func convert(s string, numRows int) string {
	l := len(s)
	if numRows == 1 {
		return s
	}
	buf := make([][]byte, min(l, numRows))
	inputs := []byte(s)

	x := 0
	isMoveDown := false
	for _, b := range inputs {
		buf[x] = append(buf[x], b)

		if x == 0 || x == numRows-1 {
			isMoveDown = !isMoveDown
		}

		if isMoveDown {
			x++
		} else {
			x--
		}
	}

	ans := make([]byte, 0, l)
	for _, b := range buf {
		ans = append(ans, b...)
	}
	return string(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
