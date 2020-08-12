package main

import "fmt"

func main() {
	// 定义的int超出范围就会报错
	var i int8
	i = 127

	fmt.Println(i)

	// i = 128 // 这里会报错
	// fmt.Println(i)

	// int, uint, rune, byte的使用
	var a int = 1<<63 - 1 // maxint
	fmt.Println(a)
	var b uint = 1<<64 - 1 // maxuint
	fmt.Println(b)
	var c byte = 'a' // 可以保存字符
	fmt.Println(c)
	var d rune = '一' // 保存unicode，可以是汉字
	fmt.Println(d)
}
