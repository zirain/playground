package pro1300

import (
	"sort"
)

func findBestValue(arr []int, target int) int {
	n := len(arr)
	sort.Ints(arr)

	// target太大，导致最大值都不能被平均值削掉，那就只能返回最大值
	if arr[n-1]*n <= target {
		return arr[n-1]
	}
	avg := target / n
	// target太小了，导致最小值都比平均值小，那么用平均值来削掉所有值
	if arr[0]*n >= target {
		// 平均值的小数部分如果大于0.5，取整数部分+1返回，整个的累加值会更接近target
		if (target%n)*2 > n {
			return avg + 1
		}
		return avg
	}

	// 走到这里说明所求的value一定在平均值和最大值之间
	// 因此比平均值小（以及等于平均值）的那些值累加得到sum
	// 然后在剩余大值所组成的新的array里面去寻找该value --- 递归
	var sum int
	tmpArr := []int{}
	// 把小于等于平均值的部分值累加
	for _, v := range arr {
		if v <= avg {
			sum += v
		} else {
			tmpArr = append(tmpArr, v)
		}
	}
	return findBestValue(tmpArr, target-sum)
}
