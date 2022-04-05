package pro954

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

func TestCanReorderDoubled(t *testing.T) {
	cases := []struct {
		arr      []int
		expected bool
	}{
		{
			[]int{3, 1, 3, 6}, false,
		},
		{
			[]int{2, 1, 2, 6}, false,
		},
		{
			[]int{4, -2, 2, -4}, true,
		},
		{
			[]int{0, 0}, true,
		},
		{
			[]int{1, 2, 4, 8}, true,
		},
		{
			[]int{-6, -3}, true,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.arr), func(t *testing.T) {
			got := canReorderDoubled(tc.arr)
			assert.Equal(t, tc.expected, got)
		})
	}
}
