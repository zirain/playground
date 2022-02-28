package pro190

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestReverseBits(t *testing.T) {
	cases := []struct {
		num, expected uint32
	}{
		{
			num:      0b00000010100101000001111010011100,
			expected: 964176192,
		},
		{
			num:      0b11111111111111111111111111111101,
			expected: 3221225471,
		},
	}
	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := reverseBits(tc.num)
			assert.Equal(t, tc.expected, got)
		})
	}
}
