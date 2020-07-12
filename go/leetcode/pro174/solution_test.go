package pro174

import "testing"

func Test1(t *testing.T) {
	dungeon := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	except := 7
	actual := calculateMinimumHP(dungeon)
	if actual != except {
		t.Fail()
	}
}
