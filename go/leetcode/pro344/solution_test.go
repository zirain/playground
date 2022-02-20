package pro344

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		s        []byte
		expected []byte
	}{
		{
			s:        []byte{'h', 'e', 'l', 'l', 'o'},
			expected: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			s:        []byte{'a'},
			expected: []byte{'a'},
		},
	}

	for _, tc := range cases {
		reverseString(tc.s)
		assert.Equal(t, tc.expected, tc.s)
	}
}
