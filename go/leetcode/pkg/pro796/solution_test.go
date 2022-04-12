package pro796

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestRotateString(t *testing.T) {
	cases := []struct {
		s    string
		goal string

		expected bool
	}{
		{"abcde", "cdeab", true},
		{"abcde", "abced", false},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := rotateString(tc.s, tc.goal)
			assert.Equal(t, tc.expected, got)
		})
	}
}
