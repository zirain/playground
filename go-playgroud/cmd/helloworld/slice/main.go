package main

import "fmt"

func SliceExpress() {
	orderLen := 5
	order := make([]uint16, 2*orderLen)
	for i := range order {
		order[i] = uint16(i)
	}

	// 扩展表达式 A[low:high:max]，只有low可以省略
	// 向A append 时，切片会扩容产生新的切片，不会覆盖原始的切片:q!
	pollorder := order[:orderLen:orderLen] // 等价于order[0:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]

	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))

	fmt.Println("pollorder = ", pollorder)
	fmt.Println("lockorder = ", lockorder)

}

func main() {
	SliceExpress()
}
