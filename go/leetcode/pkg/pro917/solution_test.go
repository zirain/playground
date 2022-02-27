package pro917

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestTeverseOnlyLetters(t *testing.T) {
	cases := []struct {
		s        string
		expected string
	}{
		{
			s:        "ab-cd",
			expected: "dc-ba",
		},
		{
			s:        "a-bC-dEf-ghIj",
			expected: "j-Ih-gfE-dCba",
		},
		{
			s:        "Test1ng-Leet=code-Q!",
			expected: "Qedo1ct-eeLg=ntse-T!",
		},
	}

	for _, tc := range cases {
		t.Run(tc.s, func(t *testing.T) {
			got := reverseOnlyLetters(tc.s)
			assert.Equal(t, tc.expected, got)
		})
	}
}
