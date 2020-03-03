package main

import(
	"fmt"
	"time"
)

// const IPv4Len = 4
// // parseIPv4 解析一个 IPv4 地址 (d.d.d.d).
// func parseIPv4(s string) IP {
//     var p [IPv4Len]byte
//     // ...
// }

//	iota 常量生成器
type Weekday int
const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func main()  {
	const (
		e  = 2.7182818
		pi = 3.1415926
	)
	
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // "1 1 2 2"
}