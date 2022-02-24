package pro540

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestSingleNonDuplicate(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			expected: 2,
		},
		{
			nums:     []int{3, 3, 7, 7, 10, 11, 11},
			expected: 10,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := singleNonDuplicate(tc.nums)
			assert.Equal(t, tc.expected, got)
		})
	}
}
