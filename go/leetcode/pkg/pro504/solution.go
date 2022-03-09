package pro504

func convertToBase7(num int) string {
	ans := make([]byte, 0)
	isNegative := false
	if num == 0 {
		return "0"
	}
	if num < 0 {
		isNegative = true
		num = -1 * num
	}
	for num != 0 {
		m := num % 7

		ans = append(ans, '0'+byte(m))
		num = num / 7

	}

	if isNegative {
		ans = append(ans, '-')
	}

	reverse(ans)
	return string(ans)
}

func reverse(input []byte) {
	inputLen := len(input)
	inputMid := inputLen / 2

	for i := 0; i < inputMid; i++ {
		j := inputLen - i - 1

		input[i], input[j] = input[j], input[i]
	}
}
