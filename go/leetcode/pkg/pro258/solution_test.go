package pro258

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestAddDigits(t *testing.T) {
	cases := []struct {
		num      int
		expected int
	}{
		{
			num:      38,
			expected: 2,
		},
		{
			num:      0,
			expected: 0,
		},
		{
			num:      111,
			expected: 3,
		},
		{
			num:      10,
			expected: 1,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := addDigits(tc.num)
			assert.Equal(t, tc.expected, got)
		})
	}
}
