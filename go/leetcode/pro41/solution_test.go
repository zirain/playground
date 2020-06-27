package pro41

import "testing"

func Test1(t *testing.T) {
	var datas = []struct {
		nums   []int
		expect int
	}{
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
	}

	for _, data := range datas {
		actual := firstMissingPositive(data.nums)
		if actual != data.expect {
			t.Errorf("firstMissingPositive(%d) = %d; expected %d", data.nums, actual, data.expect)
		}
	}
}
