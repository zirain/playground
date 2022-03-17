package pro720

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestLongestWord(t *testing.T) {
	cases := []struct {
		words    []string
		expected string
	}{
		{
			words:    []string{"w", "wo", "wor", "worl", "world"},
			expected: "world",
		},
		{
			words:    []string{"a", "banana", "app", "appl", "ap", "apply", "apple"},
			expected: "apple",
		},
		{
			words:    []string{"a", "banana", "app", "appl", "ap", "apple", "apply"},
			expected: "apple",
		},
		{
			words:    []string{"wo", "wor", "worl", "world"},
			expected: "",
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := longestWord(tc.words)
			assert.Equal(t, tc.expected, got)
		})
	}

}
