package pro977

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestSortedSquares(t *testing.T) {
	cases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			nums:     []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
	}

	for _, tc := range cases {
		got := sortedSquares(tc.nums)
		assert.Equal(t, tc.expected, got)
	}
}
