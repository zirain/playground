package pro198

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestRob(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 1, 3, 1},
			expected: 4,
		},
		{
			nums:     []int{2, 7, 9, 3, 1},
			expected: 12,
		},
		{
			nums:     []int{2, 1, 1, 2},
			expected: 4,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := rob(tc.nums)
			assert.Equal(t, tc.expected, got)
		})
	}
}
