package pro542

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestUpdateMatrix(t *testing.T) {
	cases := []struct {
		mat [][]int

		expected [][]int
	}{
		{
			mat: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 1, 0},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 1, 0},
			},
		},
		{
			mat: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1},
			},
		},
		{
			mat: [][]int{
				{0, 0, 0, 1},
				{0, 1, 0, 1},
				{0, 0, 0, 1},
				{0, 0, 0, 1},
			},
			expected: [][]int{
				{0, 0, 0, 1},
				{0, 1, 0, 1},
				{0, 0, 0, 1},
				{0, 0, 0, 1},
			},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := updateMatrix(tc.mat)
			assert.Equal(t, tc.expected, got)
		})
	}
}
