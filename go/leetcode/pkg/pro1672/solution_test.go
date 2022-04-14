package pro1672

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMaximumWealth(t *testing.T) {
	cases := []struct {
		account  [][]int
		expected int
	}{
		{
			[][]int{{1, 2, 3}, {3, 2, 1}},
			6,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := maximumWealth(tc.account)
			assert.Equal(t, tc.expected, got)
		})
	}
}
