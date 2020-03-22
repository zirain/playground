package main

import (
	"fmt"
	"sort"
)

<<<<<<< HEAD
func minIncrementForUnique(A []int) int {

	sort.Ints(A)
	ans, taken := 0, 0

	for i := 1; i < len(A); i++ {
		if A[i-1] == A[i] {
			taken++
			ans -= A[i]
		} else {
			var give = min(taken, A[i]-A[i-1]-1)
			ans += give*(give+1)/2 + give*A[i-1]
			taken -= give
		}
	}
	if len(A) > 0 {
		ans += taken*(taken+1)/2 + taken*A[len(A)-1]
	}

	return ans
}

func min(x int, y int) int {
	if x < y {
		return x
	}

	return y
}

func main() {
	fmt.Println(minIncrementForUnique([]int{1, 2, 2}))
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
=======
func longestPalindrome(s string) int {
	count:=0
	countDic := make(map[rune]int)
	for _,chr := range s{
		if _,ok:=countDic[chr]; ok{
			countDic[chr]--
			delete(countDic,chr)
			count++
		}else	{
			countDic[chr] = 0
		}
	}

	if len(countDic) > 0{
		return count*2+1
	}else{
		return count*2
	}
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
>>>>>>> ddd6444f2a9a89cf8c439d82536610f4887bcfb0
}
