package pro2100

import (
	"fmt"
	"sort"
	"testing"

	"github.com/bmizerany/assert"
)

func TestGoodDaysToRobBank(t *testing.T) {
	cases := []struct {
		security []int
		time     int
		expected []int
	}{
		{
			security: []int{5, 3, 3, 3, 5, 6, 2},
			time:     2,
			expected: []int{2, 3},
		},
		{
			security: []int{1, 1, 1, 1, 1},
			time:     0,
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			security: []int{1, 2, 3, 4, 5, 6},
			time:     2,
			expected: []int{},
		},
		{
			security: []int{1},
			time:     5,
			expected: []int{},
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v->%d", tc.security, tc.time), func(t *testing.T) {
			got := goodDaysToRobBank(tc.security, tc.time)
			sort.Ints(got)
			assert.Equal(t, tc.expected, got)
		})
	}
}
