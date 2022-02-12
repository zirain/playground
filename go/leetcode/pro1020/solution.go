package pro1020

var directions = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func numEnclaves(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if (i == 0 || i == rows-1 || j == 0 || j == cols-1) &&
				grid[i][j] == 1 {
				travelGrid(grid, i, j)
			}
		}
	}

	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				ans++
			}
		}
	}
	return ans
}

func travelGrid(grid [][]int, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 {
		return
	}

	grid[x][y] = 0
	for _, d := range directions {
		travelGrid(grid, x+d[0], y+d[1])
	}
}
