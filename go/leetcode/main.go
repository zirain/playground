package main

import "fmt"

func gameOfLife(board [][]int) {
	rows := len(board)
	cols := len(board[0])

	neighbors := []int{0, 1, -1}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			liveNeighbors := 0

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {

					if !(neighbors[i] == 0 && neighbors[j] == 0) {
						r := row + neighbors[i]
						c := col + neighbors[j]

						if (r < rows && r >= 0) && (c < cols && c >= 0) && (board[r][c] == 1 || board[r][c] == -1) {
							liveNeighbors += 1
						}
					}
				}
			}

			// 规则 1 或规则 3
			if (board[row][col] == 1) && (liveNeighbors < 2 || liveNeighbors > 3) {
				// -1 代表这个细胞过去是活的现在死了
				board[row][col] = -1
			}
			// 规则 4
			if board[row][col] == 0 && liveNeighbors == 3 {
				// 2 代表这个细胞过去是死的现在活了
				board[row][col] = 2
			}
		}
	}

	// 遍历 board 得到一次更新后的状态
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] > 0 {
				board[row][col] = 1
			} else {
				board[row][col] = 0
			}
		}
	}
}

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0}}
	gameOfLife(board)
	fmt.Println(board)
}
