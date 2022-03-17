package pro720

import (
	"sort"
)

func longestWord(words []string) string {
	sort.Strings(words)
	ans, maxLen := "", 0
	dp := make(map[string]struct{})
	for _, word := range words {
		l := len(word)
		prefix := word[:l-1]
		if _, ok := dp[prefix]; ok {
			dp[word] = struct{}{}
			if l > maxLen {
				ans, maxLen = word, l
			}
		}
	}

	return ans
}
