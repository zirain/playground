package pro762

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestCountPrimeSetBits(t *testing.T) {
	cases := []struct {
		left, right int

		expected int
	}{
		{
			6, 10, 4,
		},
		{
			10, 15, 5,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := countPrimeSetBits(tc.left, tc.right)
			assert.Equal(t, tc.expected, got)
		})
	}
}
