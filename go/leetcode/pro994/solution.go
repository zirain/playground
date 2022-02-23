package pro994

var (
	deltaX = []int{1, -1, 0, 0}
	deltaY = []int{0, 0, 1, -1}
)

func orangesRotting(grid [][]int) int {

	var (
		rows       = len(grid)
		cols       = len(grid[0])
		base       = rows * cols
		mintues    = 0
		freshCount = 0
	)
	rotted := make([]int, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				freshCount++
			} else if grid[i][j] == 2 {
				rotted = append(rotted, i*base+j)
			}
		}
	}

	if freshCount == 0 {
		return 0
	}

	for len(rotted) != 0 {
		mintues++
		rottedOranged := rotted
		rotted = make([]int, 0)

		for i := 0; i < len(rottedOranged); i++ {
			x := rottedOranged[i] / base
			y := rottedOranged[i] % base
			for n := 0; n < 4; n++ {
				newX := x + deltaX[n]
				newY := y + deltaY[n]
				cnt := 0
				if newX >= 0 && newX < rows &&
					newY >= 0 && newY < cols &&
					grid[newX][newY] == 1 {

					grid[newX][newY] = 2
					cnt++
				}

				if cnt > 0 {
					freshCount -= cnt
					rotted = append(rotted, newX*base+newY)
				}
			}
		}
	}

	if freshCount > 0 {
		return -1
	}
	return mintues - 1
}
