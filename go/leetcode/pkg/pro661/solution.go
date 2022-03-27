package pro661

var directions = [][]int{
	{1, 1}, {1, 0}, {1, -1},
	{0, 1}, {0, 0}, {0, -1},
	{-1, 1}, {-1, 0}, {-1, -1}}

func imageSmoother(img [][]int) [][]int {
	rows, cols := len(img), len(img[0])
	ans := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			ans[i][j] = calcSmoother(img, i, j)
		}
	}

	return ans
}

func calcSmoother(img [][]int, i, j int) int {
	rows, cols := len(img), len(img[0])
	total, count := 0, 0
	for _, dir := range directions {
		x, y := i+dir[0], j+dir[1]
		if x < 0 || x >= rows ||
			y < 0 || y >= cols {
			continue
		}

		total += img[x][y]
		count++
	}

	return total / count
}
