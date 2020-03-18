package main

import (
	"fmt"
)

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(rec1[2] <= rec2[0] || // left
		rec1[3] <= rec2[1] || // bottom
		rec1[0] >= rec2[2] || // right
		rec1[1] >= rec2[3]) // top
}

func main() {
	fmt.Println(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3}))
	fmt.Println(isRectangleOverlap([]int{0, 0, 1, 1}, []int{1, 0, 2, 1}))
}
