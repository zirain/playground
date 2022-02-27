package pro784

import (
	"sort"
	"testing"

	"github.com/bmizerany/assert"
)

func TestLetterCasePermutation(t *testing.T) {
	cases := []struct {
		s string

		expected []string
	}{
		{
			s:        "a1b2",
			expected: []string{"a1b2", "a1B2", "A1b2", "A1B2"},
		},
		{
			s:        "a12b",
			expected: []string{"a12b", "a12B", "A12b", "A12B"},
		},
		{
			s:        "3z4",
			expected: []string{"3Z4", "3z4"},
		},
		{
			s:        "FjkZh",
			expected: []string{"fjkzh", "fjkzH", "fjkZh", "fjkZH", "fjKzh", "fjKzH", "fjKZh", "fjKZH", "fJkzh", "fJkzH", "fJkZh", "fJkZH", "fJKzh", "fJKzH", "fJKZh", "fJKZH", "Fjkzh", "FjkzH", "FjkZh", "FjkZH", "FjKzh", "FjKzH", "FjKZh", "FjKZH", "FJkzh", "FJkzH", "FJkZh", "FJkZH", "FJKzh", "FJKzH", "FJKZh", "FJKZH"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.s, func(t *testing.T) {
			got := letterCasePermutation(tc.s)
			sort.Strings(tc.expected)
			sort.Strings(got)
			assert.Equal(t, tc.expected, got)
		})
	}
}
