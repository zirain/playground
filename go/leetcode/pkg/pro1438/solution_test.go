package pro1438

import "testing"

func TestLongestSubarray(t *testing.T) {
	type sampleData struct {
		nums   []int
		limit  int
		expect int
	}

	samples := []sampleData{
		{
			[]int{8, 2, 4, 7},
			4,
			2,
		},
		{
			[]int{10, 1, 2, 4, 7, 2},
			5,
			4,
		},
		{
			[]int{4, 2, 2, 2, 4, 4, 2, 2},
			0,
			3,
		},
	}

	for _, d := range samples {
		actual := longestSubarray(d.nums, d.limit)
		if d.expect != actual {
			t.Errorf("expect: %d actual:%d", d.expect, actual)
		}
	}
}
