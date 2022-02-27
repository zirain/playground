package pro718

import ("testing"
)

func Test1(t *testing.T) {
	A := []int{1,2,3,2,1}
	B := []int{3,2,1,4,7}
	except := 3
	actual := findLength(A,B)
	if except != actual{
		t.Fail()
	}
}