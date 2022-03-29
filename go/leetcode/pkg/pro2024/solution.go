package pro2024

func maxConsecutiveAnswers(answerKey string, k int) int {
	cnt := [26]int{}
	maxCnt, left := 0, 0
	for right, ch := range answerKey {
		cnt[ch-'A']++
		maxCnt = max(maxCnt, cnt[ch-'A'])
		if right-left+1-maxCnt > k {
			cnt[answerKey[left]-'A']--
			left++
		}
	}
	return len(answerKey) - left
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
