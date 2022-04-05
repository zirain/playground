package pro34

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestSearchRange(t *testing.T) {
	cases := []struct {
		nums   []int
		target int

		expected []int
	}{
		{
			[]int{5, 7, 7, 8, 8, 10},
			8,
			[]int{3, 4},
		},
		{
			[]int{5, 7, 7, 8, 8, 10},
			6,
			[]int{-1, -1},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := searchRange(tc.nums, tc.target)
			assert.Equal(t, tc.expected, got)
		})
	}
}
