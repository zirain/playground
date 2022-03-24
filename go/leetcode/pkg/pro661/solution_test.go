package pro661

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestImageSmoother(t *testing.T) {
	cases := []struct {
		img      [][]int
		expected [][]int
	}{
		{
			img: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			img: [][]int{
				{100, 200, 100},
				{200, 50, 200},
				{100, 200, 100},
			},
			expected: [][]int{
				{137, 141, 137},
				{141, 138, 141},
				{137, 141, 137},
			},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := imageSmoother(tc.img)
			assert.Equal(t, tc.expected, got)
		})
	}
}
