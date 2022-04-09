package pro804

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestUniqueMorseRepresentations(t *testing.T) {
	cases := []struct {
		words    []string
		expected int
	}{
		{
			words:    []string{"gin", "zen", "gig", "msg"},
			expected: 2,
		},
		{
			words:    []string{"a"},
			expected: 1,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := uniqueMorseRepresentations(tc.words)
			assert.Equal(t, tc.expected, got)
		})
	}
}
