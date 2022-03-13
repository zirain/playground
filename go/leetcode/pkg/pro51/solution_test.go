package pro51

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

func TestSolveNQueens(t *testing.T) {
	cases := []struct {
		n        int
		expected [][]string
	}{
		{
			n: 4,
			expected: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			n: 1,
			expected: [][]string{
				{"Q"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d", tc.n), func(t *testing.T) {
			got := solveNQueens(tc.n)
			assert.Equal(t, tc.expected, got)
		})
	}
}
