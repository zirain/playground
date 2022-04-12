package pro806

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNumberOfLines(t *testing.T) {
	cases := []struct {
		widths   []int
		s        string
		expected []int
	}{
		{
			widths:   []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			s:        "abcdefghijklmnopqrstuvwxyz",
			expected: []int{3, 60},
		},
		{
			widths:   []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			s:        "bbbcccdddaaa",
			expected: []int{2, 4},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := numberOfLines(tc.widths, tc.s)
			assert.Equal(t, tc.expected, got)
		})
	}
}
