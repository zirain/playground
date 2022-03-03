package pro1706

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestFindBall(t *testing.T) {
	cases := []struct {
		grid     [][]int
		expected []int
	}{
		{
			grid: [][]int{
				{1, 1, 1, -1, -1},
				{1, 1, 1, -1, -1},
				{-1, -1, -1, 1, 1},
				{1, 1, 1, 1, -1},
				{-1, -1, -1, -1, -1},
			},
			expected: []int{1, -1, -1, -1, -1},
		},
		{
			grid: [][]int{
				{-1},
			},
			expected: []int{-1},
		},
		{
			grid: [][]int{
				{-1, -1, -1},
			},
			expected: []int{-1, 0, 1},
		},
		{
			grid: [][]int{
				{1, 1, 1, 1, 1, 1},
				{-1, -1, -1, -1, -1, -1},
				{1, 1, 1, 1, 1, 1, 1},
				{-1, -1, -1, -1, -1, -1, -1},
			},
			expected: []int{0, 1, 2, 3, 4, -1},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := findBall(tc.grid)
			assert.Equal(t, tc.expected, got)
		})
	}
}
