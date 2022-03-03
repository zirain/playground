package pro189

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, l, r int) {
	for l <= r {
		nums[r], nums[l] = nums[l], nums[r]

		l++
		r--
	}
}
