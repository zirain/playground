package pro283

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMoveZereo(t *testing.T) {
	cases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			nums:     []int{0},
			expected: []int{0},
		},
		{
			nums:     []int{0, 1},
			expected: []int{1, 0},
		},
		{
			nums:     []int{1, 0},
			expected: []int{1, 0},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			moveZeroes(tc.nums)
			assert.Equal(t, tc.expected, tc.nums)
		})
	}
}
