package pro46

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestPermute(t *testing.T) {
	cases := []struct {
		nums     []int
		expected [][]int
	}{
		{
			nums:     []int{1, 2, 3},
			expected: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := permute(tc.nums)
			assert.Equal(t, tc.expected, got)
		})
	}
}
