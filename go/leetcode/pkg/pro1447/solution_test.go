package pro1447

import (
	"strconv"
	"testing"

	"github.com/bmizerany/assert"
)

func TestSimplifiedFractions(t *testing.T) {
	cases := []struct {
		n      int
		expect []string
	}{
		{1, []string{}},
		{2, []string{"1/2"}},
		{3, []string{"1/2", "1/3", "2/3"}},
		{4, []string{"1/2", "1/3", "1/4", "2/3", "3/4"}},
	}

	for _, tc := range cases {
		t.Run(strconv.Itoa(tc.n), func(t *testing.T) {
			got := simplifiedFractions(tc.n)
			assert.Equal(t, tc.expect, got)
		})
	}
}
