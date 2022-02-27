package pro120

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMinimumTotal(t *testing.T) {
	cases := []struct {
		triangle [][]int
		expected int
	}{
		{
			triangle: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			expected: 11,
		},
		{
			triangle: [][]int{
				{-10},
			},
			expected: -10,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := minimumTotal(tc.triangle)
			assert.Equal(t, tc.expected, got)
		})
	}
}
