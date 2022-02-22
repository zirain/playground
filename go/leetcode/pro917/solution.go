package pro917

func reverseOnlyLetters(s string) string {
	if len(s) <= 1 {
		return s
	}

	b := []byte(s)
	l, r := findNextLetter(b, 0, 1), findNextLetter(b, len(b)-1, -1)
	for l <= r {
		b[l], b[r] = b[r], b[l]
		l, r = findNextLetter(b, l+1, 1), findNextLetter(b, r-1, -1)
	}

	return string(b)
}

func findNextLetter(b []byte, start, added int) int {
	for i := start; i < len(b) && i >= 0; {
		chr := b[i]
		if isLetter(chr) {
			return i
		}
		i += added
	}

	if added > 0 {
		return len(b)
	}

	return -1
}

func isLetter(chr byte) bool {
	if (chr >= 'a' && chr <= 'z') ||
		(chr >= 'A' && chr <= 'Z') {
		return true
	}

	return false
}
