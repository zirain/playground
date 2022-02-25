package pro77

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestCombine(t *testing.T) {
	cases := []struct {
		n, k     int
		expected [][]int
	}{
		{
			n: 1,
			k: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			n: 2,
			k: 2,
			expected: [][]int{
				{1, 2},
			},
		},
		{
			n: 4,
			k: 2,
			expected: [][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{3, 4},
			},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := combine(tc.n, tc.k)
			assert.Equal(t, tc.expected, got)
		})
	}
}
