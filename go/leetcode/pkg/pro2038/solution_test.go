package pro2038

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestWinnerOfGame(t *testing.T) {
	cases := []struct {
		colors   string
		expected bool
	}{
		{
			colors:   "AAABABB",
			expected: true,
		},
		{
			colors:   "AA",
			expected: false,
		},
		{
			colors:   "ABBBBBBBAAA",
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.colors, func(t *testing.T) {
			got := winnerOfGame(tc.colors)
			assert.Equal(t, tc.expected, got)
		})
	}
}
