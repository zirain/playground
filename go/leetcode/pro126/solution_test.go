package pro126

import (
	"fmt"
	"testing"
)

func TestFindLadders(t *testing.T) {
	benginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	ans := findLadders(benginWord, endWord, wordList)
	fmt.Println(ans)
}

func TestFindLadders2(t *testing.T) {
	benginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log"}

	ans := findLadders(benginWord, endWord, wordList)
	fmt.Println(ans)
}

func TestFindLadders3(t *testing.T) {
	benginWord := "red"
	endWord := "tax"
	wordList := []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}

	ans := findLadders(benginWord, endWord, wordList)
	fmt.Println(ans)
}

func TestFindLadders4(t *testing.T) {
	benginWord := "qa"
	endWord := "sq"
	wordList := []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}

	ans := findLadders(benginWord, endWord, wordList)
	fmt.Println(ans)
}
