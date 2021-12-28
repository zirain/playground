package climbingstairs

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestClimbStairs(t *testing.T) {
	cases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := climbStairs(tc.n)
			assert.Equal(t, tc.expected, got)
		})
	}
}
