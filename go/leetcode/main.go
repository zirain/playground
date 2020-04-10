package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	words := make([]string, 0)
	endIndex := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			words = append(words, s[i+1:endIndex+1])
			for i >= 0 && s[i] == ' ' {
				i--
			}
			endIndex = i
		}
	}
	words = append(words, s[:endIndex+1])

	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords("     the sky is blue"))
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(reverseWords("a good   example"))
}
