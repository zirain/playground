package pro682

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestCalPoint(t *testing.T) {
	cases := []struct {
		ops      []string
		expected int
	}{
		{
			ops:      []string{"1"},
			expected: 1,
		},
		{
			ops:      []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			expected: 27,
		},
		{
			ops:      []string{"5", "2", "C", "D", "+"},
			expected: 30,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := calPoints(tc.ops)
			assert.Equal(t, tc.expected, got)
		})
	}
}
