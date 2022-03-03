package pro6

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestConvert(t *testing.T) {
	cases := []struct {
		numRows     int
		s, expected string
	}{
		{
			s:        "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			s:        "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},
		{
			s:        "A",
			numRows:  1,
			expected: "A",
		},
	}

	for _, tc := range cases {
		t.Run(tc.s, func(t *testing.T) {
			got := convert(tc.s, tc.numRows)
			assert.Equal(t, tc.expected, got)
		})
	}
}
