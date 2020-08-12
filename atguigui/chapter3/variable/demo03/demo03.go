package main

import "fmt"

// 全局变量的声明，只能用var
var x1 = 100
var x2 = 200
var name1 = "mary"

// 简介方式
var (
	x3 = 100
	x4 = 200
	x5 = "haha"
)

func main() {
	// 同时声明多个变量, 方法一，只能是同一种类型

	var n1, n2, n3 int

	fmt.Println(n1, n2, n3)

	// 同时赋值
	var m1, name, m3 = 100, "tom", 888

	fmt.Println(m1, name, m3)

	// 使用直接赋值方式
	i1, name, i3 := 200, "jack", 999

	fmt.Println(i1, name, i3)

	// 输出全局变量
	fmt.Println(x1, x3, x5)
}
