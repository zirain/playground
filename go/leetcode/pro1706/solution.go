package pro1706

type point struct {
	x, y int
}

func findBall(grid [][]int) []int {
	cols := len(grid[0])
	ans := make([]int, cols)
	for i := 0; i < cols; i++ {
		ans[i] = getAns(grid, i)
	}

	return ans
}

func getAns(grid [][]int, curCol int) int {
	rows := len(grid)
	// defer func() {
	// 	fmt.Print("\n")
	// }()

	from := point{-1, curCol}
	to := point{0, curCol}
	//fmt.Printf("%d,%d -> %d,%d", from.x, from.y, to.x, to.y)
	for to.x < rows {
		next, ok := nexthop(grid, from, to)
		//fmt.Printf(" -> %d,%d %v", next.x, next.y, ok)
		if !ok {
			return -1
		}
		if next.x == from.x && from.y == next.y {
			if next.x == rows-1 {
				return next.y
			}

			return -1
		}
		from = to
		to = next
	}

	return from.y
}

func nexthop(grid [][]int, from point, to point) (point, bool) {
	rows := len(grid)
	cols := len(grid[0])
	if to.x < 0 || to.x >= rows ||
		to.y < 0 || to.y >= cols {
		return from, false
	}

	if grid[to.x][to.y] == 1 { //挡板向右
		if from.x == to.x && from.y < to.y {
			// 向下移动
			return point{
				x: to.x + 1,
				y: to.y,
			}, true
		} else if from.x == to.x && from.y > to.y {
			// 原地不动
			return from, false
		} else {
			// 向右移动
			return point{
				x: to.x,
				y: to.y + 1,
			}, true
		}
	}

	// 挡板向左
	if from.x == to.x && from.y < to.y {
		// 原地不动
		return from, false
	} else if from.x == to.x && from.y > to.y {
		// 向下移动
		return point{
			x: to.x + 1,
			y: to.y,
		}, true
	} else {
		return point{
			x: to.x,
			y: to.y - 1,
		}, true
	}
}
