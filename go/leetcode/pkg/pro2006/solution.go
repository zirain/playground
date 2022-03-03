package pro2006

func countKDifference(nums []int, k int) int {
	var (
		ans = 0
		cnt = map[int]int{}
	)
	for _, num := range nums {
		ans += cnt[num-k] + cnt[num+k]
		cnt[num]++
	}
	return ans
}
