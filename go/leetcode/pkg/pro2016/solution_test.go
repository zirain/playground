package pro2016

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMaximumDifference(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{7, 1, 5, 4},
			expected: 4,
		},
		{
			nums:     []int{9, 4, 3, 2},
			expected: -1,
		},
		{
			nums:     []int{1, 5, 2, 10},
			expected: 9,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := maximumDifference(tc.nums)
			assert.Equal(t, tc.expected, got)
		})
	}
}
