package pro46

import "fmt"

//这里声明一个全局变量用来存储所有的排列
var result [][]int

func permute(nums []int) [][]int {

	//每次调用重置result全局变量，防止结果缓存
	result = make([][]int, 0, 2*len(nums))

	//当nums只有一个元素的情况下，直接返回即可
	if len(nums) == 1 {
		result = append(result, nums)
		return result
	}

	//声明一个arr变量，用来存储路径
	arr := make([]int, 0, len(nums))

	arrange(nums, arr)

	return result
}

func arrange(nums []int, arr []int) {

	//当nums长度为0，选择列表为空，路径选择完毕，返回即可
	if len(nums) == 0 {
		return
	}

	//循环当前的选择列表
	for k, v := range nums {

		//选取一个元素，例如选1，则使用newArr来存储更新后的选择列表
		arr = append(arr, v)

		newArr := make([]int, len(nums)-1)

		copy(newArr[:k], nums[:k])

		if k < len(nums)-1 {
			copy(newArr[k:], nums[k+1:])
		}

		//递归调用，将更新后的选择列表和存储路径的arr传入
		arrange(newArr, arr)

		//当arr的长度和容量相等，说明路径已经选择完毕
		//将此条路径记录到结果中
		if len(arr) == cap(arr) {
			path := make([]int, len(arr))
			copy(path, arr)
			result = append(result, path)
		}

		/**
		返回上一个路口重新做选择
		例如，当arr为1开头的时候，下一个路口有2，3两种选择
		当选择完毕，将路径记录之后，需要返回重新选择
		例如从1->2 返回到 1->3的结果
		**/
		arr = arr[:len(arr)-1]
	}
}

func render(input [][]int) {
	fmt.Println("[")
	for i := 0; i < len(input); i++ {
		fmt.Println(" ", input[i])
	}
	fmt.Println("]")
}
