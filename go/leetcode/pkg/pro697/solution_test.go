package pro697

import "testing"

func TestFindShortestSubArray(t *testing.T) {
	type sampleData struct {
		nums   []int
		expect int
	}

	samples := []sampleData{
		{
			[]int{1, 2, 2, 3, 1},
			2,
		},
		{
			[]int{1, 2, 2, 3, 1, 4, 2},
			6,
		},
	}

	for _, d := range samples {
		actual := findShortestSubArray(d.nums)
		if d.expect != actual {
			t.Fail()
		}
	}
}
