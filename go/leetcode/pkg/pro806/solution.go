package pro806

func numberOfLines(widths []int, s string) []int {
	rows, rowLen := 1, 100
	for i := 0; i < len(s); i++ {
		w := widths[s[i]-'a']
		if rowLen >= w {
			rowLen -= w
		} else {
			rows++
			rowLen = 100 - w
		}
	}

	return []int{rows, 100 - rowLen}
}
