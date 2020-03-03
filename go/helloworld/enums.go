package main

import(
	"fmt"
)

type Weapon int

const (
	Arrow Weapon = iota
	Shuriken
	SniperRifle
	Rifle
	Blower
)

const (
    FlagNone = 1 << iota
    FlagRed
    FlagGreen
    FlagBlue
)

func main()  {
	// 输出所有枚举值
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
	// 使用枚举类型并赋初值
	var weapon Weapon = Blower
	fmt.Println(weapon)

	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("%b %b %b\n", FlagRed, FlagGreen, FlagBlue)
}