package pro693

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestHasAlternatingBits(t *testing.T) {
	cases := []struct {
		n        int
		expected bool
	}{
		{
			n:        5,
			expected: true,
		},
		{
			n:        7,
			expected: false,
		},
		{
			n:        11,
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := hasAlternatingBits(tc.n)
			assert.Equal(t, tc.expected, got)
		})
	}
}
