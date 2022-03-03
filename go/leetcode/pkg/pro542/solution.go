package pro542

func updateMatrix(mat [][]int) [][]int {
	rows := len(mat)
	cols := len(mat[0])
	ans := initMatrix(rows, cols)

	// mat[i][j] == 0 时, ans[i][j] = 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 0 {
				ans[i][j] = 0
			}
		}
	}

	// 只有 水平向左移动 和 竖直向上移动
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if ans[i][j] == 0 {
				continue
			}

			if i-1 >= 0 {
				ans[i][j] = min(ans[i][j], ans[i-1][j]+1)
			}

			if j-1 >= 0 {
				ans[i][j] = min(ans[i][j], ans[i][j-1]+1)
			}
		}
	}

	// 只有 水平向右移动 和 竖直向下移动，注意动态规划的计算顺序
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if ans[i][j] == 0 {
				continue
			}

			if i+1 < rows {
				ans[i][j] = min(ans[i][j], ans[i+1][j]+1)
			}
			if j+1 < cols {
				ans[i][j] = min(ans[i][j], ans[i][j+1]+1)
			}
		}
	}

	return ans
}

func initMatrix(rows, cols int) [][]int {
	ans := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			ans[i][j] = rows * cols // 本题中 rows*cols已经足够大
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
