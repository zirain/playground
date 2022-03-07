package pro504

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestConvertToBase7(t *testing.T) {
	cases := []struct {
		num      int
		expected string
	}{
		{
			num:      100,
			expected: "202",
		},
		{
			num:      -7,
			expected: "-10",
		},
		{
			num:      0,
			expected: "0",
		},
		{
			num:      9,
			expected: "12",
		},
		{
			num:      1,
			expected: "1",
		},
	}
	for _, tc := range cases {
		got := convertToBase7(tc.num)
		assert.Equal(t, tc.expected, got)
	}
}
