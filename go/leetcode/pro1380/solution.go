package pro1380

func luckyNumbers(matrix [][]int) []int {
	ans := make([]int, 0)

	rows := len(matrix)
	for i := 0; i < rows; i++ {
		min, minIdex := min(matrix[i])
		if isMaxInCol(matrix, i, minIdex) {
			ans = append(ans, min)
		}
	}

	return ans
}

func isMaxInCol(matrix [][]int, row, col int) bool {
	n := matrix[row][col]
	for i := 0; i < len(matrix); i++ {
		if matrix[i][col] > n {
			return false
		}
	}

	return true
}

func min(nums []int) (int, int) {
	m, mIndex := nums[0], 0
	for i, n := range nums {
		if n < m {
			m, mIndex = n, i
		}
	}

	return m, mIndex
}
