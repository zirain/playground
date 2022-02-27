package pro1791

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestFindCenter(t *testing.T) {
	cases := []struct {
		edges    [][]int
		expected int
	}{
		{
			edges:    [][]int{{1, 2}, {2, 3}, {2, 4}},
			expected: 2,
		},
		{
			edges:    [][]int{{1, 2}, {1, 3}, {1, 4}, {5, 1}},
			expected: 1,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := findCenter(tc.edges)
			assert.Equal(t, tc.expected, got)
		})
	}
}
