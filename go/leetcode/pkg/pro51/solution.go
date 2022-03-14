package pro51

var ans [][]string

func solveNQueens(n int) [][]string {
	ans = make([][]string, 0)
	board := initBoard(n)
	backtrack(board, 0)
	return ans
}

func backtrack(board [][]byte, row int) {
	if row == len(board) {
		ans = append(ans, boardToString(board))
		return
	}

	n := len(board[row])
	for col := 0; col < n; col++ {
		if !isValid(board, row, col) {
			continue
		}

		board[row][col] = 'Q'
		backtrack(board, row+1)
		board[row][col] = '.'
	}
}

func isValid(board [][]byte, row, col int) bool {
	n := len(board)
	// 检查列是否有皇后互相冲突
	for i := 0; i < n; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 检查右上方是否有皇后互相冲突
	for i, j := row-1, col+1; i >= 0 && j < n; {
		if board[i][j] == 'Q' {
			return false
		}

		i--
		j++
	}

	// 检查左上方是否有皇后互相冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if board[i][j] == 'Q' {
			return false
		}

		i--
		j--
	}
	return true
}

func boardToString(board [][]byte) []string {
	ans := make([]string, len(board))
	for i := 0; i < len(board); i++ {
		ans[i] = string(board[i])
	}
	return ans
}

func initBoard(n int) [][]byte {
	ans := make([][]byte, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			ans[i][j] = '.'
		}
	}

	return ans
}
