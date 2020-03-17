package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	rows := len(grid)

	cols := len(grid[0])

	maxArea := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				continue
			}

			area := dfs(grid, rows, cols, i, j, 0)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func dfs(grid [][]int, rows int, cols int, i int, j int, area int) int {
	dirX := []int{1, -1, 0, 0}
	dirY := []int{0, 0, 1, -1}

	area++
	grid[i][j] = 0

	for n := 0; n < len(dirX); n++ {
		newI := i + dirX[n]
		newJ := j + dirY[n]

		if newI >= 0 &&
			newI < rows &&
			newJ >= 0 &&
			newJ < cols &&
			grid[newI][newJ] == 1 {

			area = dfs(grid, rows, cols, newI, newJ, area)
		}
	}

	return area
}

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	maxArea := maxAreaOfIsland(grid)
	fmt.Println("Max area of isliad is ", maxArea)
}
