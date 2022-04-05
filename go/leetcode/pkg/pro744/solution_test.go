package pro744

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNextGreatestLetter(t *testing.T) {
	cases := []struct {
		letters  []byte
		target   byte
		expected byte
	}{
		{
			letters:  []byte{'c', 'f', 'j'},
			target:   'a',
			expected: 'c',
		},
		{
			letters:  []byte{'c', 'f', 'j'},
			target:   'c',
			expected: 'f',
		},
		{
			letters:  []byte{'c', 'f', 'j'},
			target:   'd',
			expected: 'f',
		},
		{
			letters:  []byte{'c', 'f', 'j'},
			target:   'l',
			expected: 'c',
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := nextGreatestLetter(tc.letters, tc.target)
			assert.Equal(t, tc.expected, got)
		})
	}
}
