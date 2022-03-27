package pro2039

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNetworkBecomesIdle(t *testing.T) {
	cases := []struct {
		edges    [][]int
		patience []int
		expected int
	}{
		{
			edges: [][]int{
				{0, 1},
				{1, 2},
			},
			patience: []int{0, 2, 1},
			expected: 8,
		},
		{
			edges: [][]int{
				{0, 1},
				{0, 2},
				{1, 2},
			},
			patience: []int{0, 10, 10},
			expected: 3,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := networkBecomesIdle(tc.edges, tc.patience)
			assert.Equal(t, tc.expected, got)
		})
	}
}
