package pro55

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestCanJump(t *testing.T) {
	cases := []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := canJump(tc.nums)
			assert.Equal(t, tc.expected, got)
		})
	}
}
