package main

import (
	"fmt"
	"unsafe"
)

func main() {
	n := 100

	fmt.Printf("n 的类型是：%T \n", n)

	// 查看某个变量占用字节大小和数据类型
	var n2 int64 = 10
	fmt.Printf("n2 的类型是：%T \nn2 占用空间字节数是：%d", n2, unsafe.Sizeof(n2))
}
