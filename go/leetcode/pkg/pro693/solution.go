package pro693

func hasAlternatingBits(n int) bool {
	prev := -1
	for n > 0 {
		t := n & 1
		if t == prev {
			return false
		} else {
			prev = t
		}
		n = n >> 1
	}

	return true
}
