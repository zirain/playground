package pro1748

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_SumOfUnique(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 2, 3, 2},
			expected: 4,
		},
		{
			nums:     []int{1, 1, 1, 1},
			expected: 0,
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := sumOfUnique(tc.nums)
			assert.Equal(t, tc.expected, got)
		})
	}
}
