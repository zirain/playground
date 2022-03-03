package pro2006

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestCountKDifference(t *testing.T) {
	cases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{1, 2, 2, 1},
			k:        1,
			expected: 4,
		},
		{
			nums:     []int{1, 3},
			k:        3,
			expected: 0,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := countKDifference(tc.nums, tc.k)
			assert.Equal(t, tc.expected, got)
		})
	}
}
