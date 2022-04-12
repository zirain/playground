package pro704

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
		{
			nums:     []int{5},
			target:   5,
			expected: 0,
		},
		{
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   13,
			expected: -1,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := search(tc.nums, tc.target)
			assert.Equal(t, tc.expected, got)
		})
	}
}
