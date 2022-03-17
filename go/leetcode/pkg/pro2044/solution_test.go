package pro2044

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestCountMaxOrSubsets(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{3, 1},
			expected: 2,
		},
		{
			nums:     []int{2, 2, 2},
			expected: 7,
		},
		{
			nums:     []int{3, 2, 1, 5},
			expected: 6,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := countMaxOrSubsets(tc.nums)
			assert.Equal(t, tc.expected, got)
		})
	}
}
