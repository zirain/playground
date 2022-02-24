package pro1380

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestLuckyNumbers(t *testing.T) {
	cases := []struct {
		matrix   [][]int
		expected []int
	}{
		{
			matrix: [][]int{
				{3, 7, 8},
				{9, 11, 13},
				{15, 16, 17},
			},
			expected: []int{15},
		},
		{
			matrix: [][]int{
				{1, 10, 4, 2},
				{9, 3, 8, 7},
				{15, 16, 17, 12},
			},
			expected: []int{12},
		},
		{
			matrix: [][]int{
				{7, 8},
				{1, 2},
			},
			expected: []int{7},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := luckyNumbers(tc.matrix)
			assert.Equal(t, tc.expected, got)
		})
	}
}
