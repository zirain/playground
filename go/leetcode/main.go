package main

func numRookCaptures(board [][]byte) int {
	r := 0
	rx, ry := getRootPosition(board)
	directions := []byte{'u', 'd', 'l', 'r'}
	for x, y, i := rx, ry, 0; i < 4; {
		var d = directions[i]
		switch d {
		case 'u':
			x--
		case 'd':
			x++
		case 'l':
			y--
		case 'r':
			y++
		}

		if x >= 0 && y >= 0 && x < len(board) && y < len(board[0]) {
			chess := "" + string(board[x][y])
			if chess != "." {
				if chess == "p" {
					r++
				}
				i++
				x = rx
				y = ry
			}

		} else {
			i++
			x = rx
			y = ry
		}
	}

	return r
}

func getRootPosition(board [][]byte) (int, int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'R' {
				return i, j
			}
		}
	}
	return -1, -1
}

func main() {
}
