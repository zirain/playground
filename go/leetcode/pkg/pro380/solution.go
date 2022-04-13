package pro380

import (
	"math/rand"
)

type RandomizedSet struct {
	nums []int
	set  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: make([]int, 0),
		set:  make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}
	this.set[val] = len(this.nums)
	this.nums = append(this.nums, val)

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.set[val]
	if !ok {
		return false
	}
	last := len(this.nums) - 1
	this.nums[idx] = this.nums[last]
	this.set[this.nums[idx]] = idx
	this.nums = this.nums[:last]
	delete(this.set, val)
	return true

}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
