package pro994

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestOrangesRotting(t *testing.T) {
	cases := []struct {
		mat [][]int

		expected int
	}{
		{
			mat: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 4,
		},
		{
			mat: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			expected: -1,
		},
		{
			mat: [][]int{
				{0, 2},
			},
			expected: 0,
		},
		{
			mat: [][]int{
				{0, 0},
			},
			expected: 0,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := orangesRotting(tc.mat)
			assert.Equal(t, tc.expected, got)
		})
	}
}
