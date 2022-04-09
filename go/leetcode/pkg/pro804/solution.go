package pro804

var morses = []string{
	".-", "-...", "-.-.", "-..", ".", "..-.", "--.",
	"....", "..", ".---", "-.-", ".-..", "--", "-.",
	"---", ".--.", "--.-", ".-.", "...", "-", "..-",
	"...-", ".--", "-..-", "-.--", "--..",
}

func uniqueMorseRepresentations(words []string) int {
	res := make(map[string]struct{})
	for _, w := range words {
		m := morse(w)
		res[m] = struct{}{}
	}
	return len(res)
}

func morse(word string) string {
	ans := ""
	for i := 0; i < len(word); i++ {
		ans += morses[word[i]-'a']
	}
	return ans
}
