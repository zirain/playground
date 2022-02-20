package pro344

func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l <= r {
		s[r], s[l] = s[l], s[r]

		l++
		r--
	}
}
