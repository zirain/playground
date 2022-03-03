package pro733

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestFloodFill(t *testing.T) {
	cases := []struct {
		image     [][]int
		x, y      int
		newClolor int
		expected  [][]int
	}{
		{
			image: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			x:         1,
			y:         1,
			newClolor: 2,
			expected: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		{
			image: [][]int{
				{1, 1, 1, 0},
				{1, 1, 0, 0},
				{1, 0, 1, 0},
				{0, 0, 0, 0},
			},
			x:         1,
			y:         1,
			newClolor: 2,
			expected: [][]int{
				{2, 2, 2, 0},
				{2, 2, 0, 0},
				{2, 0, 1, 0},
				{0, 0, 0, 0},
			},
		},
		{
			image: [][]int{
				{0, 0, 0},
				{0, 1, 2},
			},
			x:         1,
			y:         1,
			newClolor: 1,
			expected: [][]int{
				{0, 0, 0},
				{0, 1, 2},
			},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := floodFill(tc.image, tc.x, tc.y, tc.newClolor)
			assert.Equal(t, tc.expected, got)
		})
	}
}
