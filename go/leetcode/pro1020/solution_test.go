package pro1020

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNumEnclaves(t *testing.T) {
	cases := []struct {
		grid     [][]int
		expected int
	}{
		{
			grid: [][]int{
				{0, 0, 0, 0},
				{1, 0, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
			expected: 3,
		},
		{
			grid: [][]int{
				{0, 1, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			expected: 0,
		},
		{
			grid: [][]int{
				{0, 1, 0, 0, 1},
				{0, 0, 1, 1, 0},
				{0, 0, 1, 1, 0},
				{0, 0, 0, 0, 1},
			},
			expected: 4,
		},
		{
			grid: [][]int{
				{0, 0, 0, 0},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			expected: 4,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := numEnclaves(tc.grid)
			assert.Equal(t, tc.expected, got)
		})
	}
}
