package pro717

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestIsOneBitCharacter(t *testing.T) {
	cases := []struct {
		bits     []int
		expected bool
	}{
		{
			bits:     []int{1, 1, 1, 0},
			expected: false,
		},
		{
			bits:     []int{1, 0, 0},
			expected: true,
		},
		{
			bits:     []int{1, 1, 0, 0},
			expected: true,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := isOneBitCharacter(tc.bits)
			assert.Equal(t, tc.expected, got)
		})
	}
}
