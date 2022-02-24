package pro577

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestReverseWords(t *testing.T) {
	cases := []struct {
		s        string
		expected string
	}{
		{
			s:        "ab",
			expected: "ba",
		},
		{
			s:        "Let's take LeetCode contest",
			expected: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			s:        "God Ding",
			expected: "doG gniD",
		},
	}
	for _, tc := range cases {
		t.Run(tc.s, func(t *testing.T) {
			got := reverseWords(tc.s)
			assert.Equal(t, tc.expected, got)
		})
	}
}
