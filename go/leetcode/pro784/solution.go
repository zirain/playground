package pro784

var ans = make([]string, 0)

func letterCasePermutation(s string) []string {
	ans = make([]string, 0)
	arrage([]byte(s), 0)
	return ans
}

func arrage(b []byte, l int) {
	if l >= len(b) {
		ans = append(ans, string(b))
	}

	for i := l; i < len(b); i++ {
		if i == len(b)-1 {
			ans = append(ans, string(b))
		}

		if !isLetter(b[i]) {
			continue
		}

		newBytes := casePermutation(b, i)
		arrage(newBytes, i+1)
	}
}

func casePermutation(b []byte, i int) []byte {
	bf := make([]byte, len(b))
	copy(bf, b)
	bf[i] = caseLetter(bf[i])
	return bf
}

func isLetter(b byte) bool {
	if (b >= 'A' && b <= 'Z') ||
		(b >= 'a' && b <= 'z') {
		return true
	}

	return false
}

const delta = 'a' - 'A'

func caseLetter(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + delta
	}

	return b - delta
}
