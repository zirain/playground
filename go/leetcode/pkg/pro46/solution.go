package pro46

//这里声明一个全局变量用来存储所有的排列
var ans [][]int

func permute(nums []int) [][]int {
	ans = make([][]int, 0, 2*len(nums))
	track := make([]int, 0)
	used := make([]bool, len(nums))

	backtrack(nums, track, used)

	return ans
}

func backtrack(nums []int, track []int, used []bool) {
	if len(track) == len(nums) {
		ans = append(ans, track)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		track = append(track, nums[i])
		used[i] = true
		backtrack(nums, track, used)
		track = track[:len(track)-1]
		used[i] = false
	}
}
