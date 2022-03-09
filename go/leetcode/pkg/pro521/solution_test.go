package pro521

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestFindLUSlength(t *testing.T) {
	cases := []struct {
		a, b     string
		expected int
	}{
		{
			a:        "aba",
			b:        "cdc",
			expected: 3,
		},
		{
			a:        "aaa",
			b:        "bbb",
			expected: 3,
		},
		{
			a:        "aaa",
			b:        "aaa",
			expected: -1,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := findLUSlength(tc.a, tc.b)
			assert.Equal(t, tc.expected, got)
		})
	}
}
