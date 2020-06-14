package pro1300

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	arr := []int{4, 9, 3}
	target := 10
	fmt.Println(findBestValue(arr, target))
}

func Test2(t *testing.T) {
	arr := []int{2, 3, 5}
	target := 10
	fmt.Println(findBestValue(arr, target))
}

func Test3(t *testing.T) {
	arr := []int{60864, 25176, 27249, 21296, 20204}
	target := 56803
	fmt.Println(findBestValue(arr, target))
}
