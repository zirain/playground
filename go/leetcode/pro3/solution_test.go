package pro3

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		s        string
		expected int
	}{
		{
			s:        "abcabcbb",
			expected: 3,
		},
		{
			s:        "bbbbb",
			expected: 1,
		},
		{
			s:        "pwwkew",
			expected: 3,
		},
		{
			s:        " ",
			expected: 1,
		},
		{
			s:        "aab",
			expected: 2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.s, func(t *testing.T) {
			got := lengthOfLongestSubstring(tc.s)
			assert.Equal(t, tc.expected, got)
		})
	}
}
