package pro537

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestComplexNumberMultiply(t *testing.T) {
	cases := []struct {
		num1, num2 string
		expected   string
	}{
		{
			num1:     "1+1i",
			num2:     "1+1i",
			expected: "0+2i",
		},
		{
			num1:     "1+-1i",
			num2:     "1+-1i",
			expected: "0+-2i",
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := complexNumberMultiply(tc.num1, tc.num2)
			assert.Equal(t, tc.expected, got)
		})

	}
}
