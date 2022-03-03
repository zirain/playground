package pro1984

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMinimumDifference(t *testing.T) {
	cases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{90},
			k:        1,
			expected: 0,
		},
		{
			nums:     []int{9, 4, 1, 7},
			k:        2,
			expected: 2,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := minimumDifference(tc.nums, tc.k)
			assert.Equal(t, tc.expected, got)
		})
	}
}
