package pro728

func selfDividingNumbers(left int, right int) []int {
	ans := make([]int, 0, right-left+1)
	for n := left; n <= right; n++ {
		if isSelfDividing(n) {
			ans = append(ans, n)
		}
	}

	return ans
}

func isSelfDividing(num int) bool {
	original := num
	divideNum := 0
	for num != 0 {
		divideNum = num % 10
		if divideNum == 0 || original%divideNum != 0 {
			return false
		}

		num /= 10
	}

	return true
}
