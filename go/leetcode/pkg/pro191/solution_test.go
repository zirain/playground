package pro191

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestHammingWeight(t *testing.T) {
	cases := []struct {
		num      uint32
		expected int
	}{
		{
			num:      0b00000000000000000000000000001011,
			expected: 3,
		},
		{
			num:      0b00000000000000000000000010000000,
			expected: 1,
		},
		{
			num:      0b11111111111111111111111111111101,
			expected: 31,
		},
		{
			num:      0b11111111111111111111111111111100,
			expected: 30,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := hammingWeight(tc.num)
			assert.Equal(t, tc.expected, got)
		})
	}
}
