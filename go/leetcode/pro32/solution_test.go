package pro32

import "testing"

func Test1(t *testing.T) {
	var datas = []struct {
		input  string
		expect int
	}{
		{"(()", 2},
		{")()())", 4},
		{"()(()", 2},
		{"()(())", 6},
	}
	for _, data := range datas {
		actual := longestValidParentheses(data.input)
		if actual != data.expect {
			t.Errorf("longestValidParentheses(%s) = %d; expected %d", data.input, actual, data.expect)
		}
	}
}
