package pro136

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			nums:     []int{4, 1, 1, 2, 2},
			expected: 4,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := singleNumber(tc.nums)
			assert.Equal(t, tc.expected, got)
		})
	}
}
