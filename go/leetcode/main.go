package main

import (
	"fmt"
)

func countCharacters(words []string, chars string) int {
	charsCounts := countChars(chars)

	wordLen := 0
	for _, val := range words {
		counts := charsCounts
		if canStringBuilt(val, counts) {
			wordLen += len(val)
		}
	}
	return wordLen
}

func canStringBuilt(chars string, charCounts [26]int) bool {

	for _, val := range chars {
		idx := val - 'a'
		if charCounts[idx] > 0 {
			charCounts[idx]--
		} else {
			return false
		}
	}

	return true
}

func countChars(chars string) [26]int {
	counts := [26]int{}
	for _, val := range chars {
		counts[val-'a']++
	}

	return counts
}

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	fmt.Println(countCharacters(words, chars))
}
