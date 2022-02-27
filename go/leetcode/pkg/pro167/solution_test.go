package pro167

import (
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
			expected: []int{1, 2},
		},
		{
			nums:     []int{2, 3, 4},
			target:   6,
			expected: []int{1, 3},
		},
		{
			nums:     []int{2, 3, 4},
			target:   5,
			expected: []int{1, 2},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := twoSum(tc.nums, tc.target)
			assert.Equal(t, tc.expected, got)
		})
	}
}
