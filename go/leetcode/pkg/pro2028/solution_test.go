package pro2028

import (
	"sort"
	"testing"

	"github.com/bmizerany/assert"
)

func TestMissingRolls(t *testing.T) {
	cases := []struct {
		rolls []int
		mean  int
		n     int

		expected []int
	}{
		{
			[]int{3, 2, 4, 3},
			4,
			2,
			[]int{6, 6},
		},
		{
			[]int{1, 5, 6},
			3,
			4,
			[]int{2, 2, 2, 3},
		},
		{
			[]int{1, 2, 3, 4},
			6,
			4,
			[]int{},
		},
		{
			[]int{1},
			3,
			1,
			[]int{5},
		},
		{
			[]int{2},
			1,
			1,
			[]int{},
		},
		{
			[]int{4, 2, 2, 5, 4, 5, 4, 5, 3, 3, 6, 1, 2, 4, 2, 1, 6, 5, 4, 2, 3, 4, 2, 3, 3, 5, 4, 1, 4, 4, 5, 3, 6, 1, 5, 2, 3, 3, 6, 1, 6, 4, 1, 3},
			2,
			53,
			[]int{},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := missingRolls(tc.rolls, tc.mean, tc.n)
			sort.Ints(got)
			assert.Equal(t, tc.expected, got)
		})
	}
}
