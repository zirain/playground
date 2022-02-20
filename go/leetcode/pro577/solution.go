package pro577

func reverseWords(s string) string {
	b := []byte(s)

	l, r := 0, findNextSpace(b, 0)
	for l < len(b) {
		reverse(b, l, r-1)
		l = r + 1
		r = findNextSpace(b, l)
	}

	return string(b)
}

func reverse(b []byte, l, r int) {
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}

func findNextSpace(b []byte, start int) int {
	for i := start; i < len(b); i++ {
		if b[i] == ' ' {
			return i
		}
	}
	return len(b)
}
