package pro733

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	cur := image[sr][sc]
	if cur != newColor {
		bfs(image, sr, sc, image[sr][sc], newColor)
	}

	return image
}

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func bfs(image [][]int, x, y int, target, newColor int) {
	rows := len(image)
	cols := len(image[0])

	if x < 0 || x >= rows ||
		y < 0 || y >= cols ||
		image[x][y] != target {
		return
	}

	image[x][y] = newColor
	for _, d := range dirs {
		newX := x + d[0]
		newY := y + d[1]
		bfs(image, newX, newY, target, newColor)
	}
}
