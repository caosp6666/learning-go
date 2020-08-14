package main

import "fmt"

// func main() {
// 	// str := 'a' + '1' // 单引号在go语言中表示golang中的rune(int32)类型，这里会输出146，即a和1的ASCII求和
// 	// str := "abc" + "123" // 成功连接
// 	// str := '123' + "abc" // 前面是rune类型，不能多于一个字符
// 	fmt.Printf("abc%d", 123)
// 	// fmt.Println(str)
// }

const (
	x, i = iota, iota
	_, _
	y, r
	z = "zz"
	k
	p = iota
)

// 间断的相当于跟上面一样的定义，iota表示常量计数

func main() {
	fmt.Println(x, y, z, k, p, i, r)
}
