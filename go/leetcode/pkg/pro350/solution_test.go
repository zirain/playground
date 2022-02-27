package pro350

import (
	"sort"
	"testing"
)

func Test1(t *testing.T) {
	var datas = []struct {
		nums1  []int
		nums2  []int
		except []int
	}{
		{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}, except: []int{2, 2}},
	}

	for _, data := range datas {
		actual := intersect(data.nums1, data.nums2)
		if !equal(actual, data.except) {
			t.Fail()
		}
	}
}

func equal(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}

	return true
}
