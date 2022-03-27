package pro1

import (
	"sort"
	"testing"

	"github.com/bmizerany/assert"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := twoSum(tc.nums, tc.target)
			sort.Ints(got)
			assert.Equal(t, tc.expected, got)
		})
	}
}
