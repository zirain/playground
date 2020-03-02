package main

import(
	"fmt"
)

func main()	{
	var str= "Hello\nWorld!"
	fmt.Println(str)

	//字符串拼接符“+”
	var str1 = "hello"
	str1 += " world!"
	fmt.Println(str1)

	//定义多行字符串
	const str2 = `第一行
第二行
第三行
\r\n
`
	fmt.Println(str2)

	//
	var ch int = '\u0041'
	var ch2 int = '\u03B2' //4字节
	var ch3 int = '\U00101234' //8字节
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
}