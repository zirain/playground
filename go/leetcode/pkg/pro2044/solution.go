package pro2044

func countMaxOrSubsets(nums []int) int {
	ans, maxOr := 0, 0

	var dfs func(pos int, or int)
	dfs = func(pos int, or int) {
		if pos == len(nums) {
			if or > maxOr {
				maxOr = or
				ans = 1
			} else if or == maxOr {
				ans++
			}

			return
		}
		dfs(pos+1, nums[pos]|or)
		dfs(pos+1, or)
	}

	dfs(0, 0)
	return ans
}
