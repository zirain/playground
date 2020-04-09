package main

import "fmt"

func generateParenthesis(n int) []string {
	output := make([]string, 0)
	if n < 0 {
		return output
	}

	backtrace("", 0, 0, n, &output)
	return output
}

func backtrace(parenthesis string, lefts int, rights int, max int, output *[]string) {
	if len(parenthesis) == (max * 2) {
		*output = append(*output, parenthesis)
	}

	if lefts < max {
		backtrace(parenthesis+"(", lefts+1, rights, max, output)
	}

	if rights < lefts {
		backtrace(parenthesis+")", lefts, rights+1, max, output)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
}
