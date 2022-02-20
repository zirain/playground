package pro189

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        4,
			expected: []int{4, 5, 6, 7, 1, 2, 3},
		},
		{
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
		{
			nums:     []int{-1},
			k:        2,
			expected: []int{-1},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			rotate(tc.nums, tc.k)
			assert.Equal(t, tc.expected, tc.nums)
		})
	}
}
