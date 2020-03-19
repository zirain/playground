package main

import (
	"fmt"
)

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
}
