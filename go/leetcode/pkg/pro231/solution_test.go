package pro231

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

func TestIsPowerOfTwo(t *testing.T) {
	cases := []struct {
		n        int
		expected bool
	}{
		{
			n:        1,
			expected: true,
		},
		{
			n:        16,
			expected: true,
		},
		{
			n:        3,
			expected: false,
		},
		{
			n:        6,
			expected: false,
		},
		{
			n:        4,
			expected: true,
		},
		{
			n:        5,
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d", tc.n), func(t *testing.T) {
			got := isPowerOfTwo(tc.n)
			assert.Equal(t, tc.expected, got)
		})
	}
}
