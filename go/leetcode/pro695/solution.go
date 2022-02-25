package pro695

func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			area := bfs(grid, i, j, 0)
			ans = max(ans, area)
		}
	}

	return ans
}

var dirs = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func bfs(grid [][]int, x, y int, area int) int {
	rows := len(grid)
	cols := len(grid[0])
	if x < 0 || x >= rows ||
		y < 0 || y >= cols ||
		grid[x][y] != 1 {
		return area
	}

	grid[x][y] = 0
	area++

	ans := area
	for i := 0; i < len(dirs); i++ {
		nextX, nextY := x+dirs[i][0], y+dirs[i][1]
		nextArea := bfs(grid, nextX, nextY, ans)
		ans = max(ans, nextArea)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
