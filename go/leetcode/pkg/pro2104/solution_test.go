package pro2104

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestSubArrayRanges(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int64
	}{
		{
			nums:     []int{1, 2, 3},
			expected: 4,
		},
		{
			nums:     []int{1, 3, 3},
			expected: 4,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := subArrayRanges(tc.nums)
			assert.Equal(t, tc.expected, got)
		})
	}
}
