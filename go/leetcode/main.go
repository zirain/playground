package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

func maxDistance(grid [][]int) int {
	var (
		count int
		stack []*point
	)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				stack = append(stack, &point{x: i, y: j})
			}
		}
	}

	dXs := []int{1, -1, 0, 0}
	dYs := []int{0, 0, 1, -1}

	for len(stack) != 0 {
		temp := stack
		stack = []*point{}
		for _, val := range temp {
			i := val.x
			j := val.y

			for n := 0; n < 4; n++ {
				newI := i + dXs[n]
				newJ := j + dYs[n]

				if newI >= 0 && newI < len(grid) && newJ >= 0 && newJ < len(grid[0]) && grid[newI][newJ] == 0 {
					grid[newI][newJ] = 1
					stack = append(stack, &point{x: newI, y: newJ})
				}
			}
		}

		if len(stack) != 0 {
			count++
		}
	}

	if count == 0 {
		return -1
	}
	return count
}

func main() {
	fmt.Println(maxDistance([][]int{[]int{1, 0, 1}, []int{0, 0, 0}, []int{1, 0, 1}}))
	fmt.Println(maxDistance([][]int{[]int{1, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}))
}
