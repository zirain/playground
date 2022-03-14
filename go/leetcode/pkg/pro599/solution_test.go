package pro599

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestFindRestaurant(t *testing.T) {
	cases := []struct {
		list1    []string
		list2    []string
		expected []string
	}{
		{
			list1:    []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			list2:    []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			expected: []string{"Shogun"},
		},
		{
			list1:    []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			list2:    []string{"KFC", "Shogun", "Burger King"},
			expected: []string{"Shogun"},
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			got := findRestaurant(tc.list1, tc.list2)
			assert.Equal(t, tc.expected, got)
		})
	}

}
