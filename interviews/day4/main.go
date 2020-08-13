package main

import "fmt"

// func main() {
// 	s1 := []int{1, 2, 3}
// 	s2 := []int{4, 5}
// 	s1 = append(s1, s2...) // 将一个切片添加到另一个切片后面的方法
// 	fmt.Println(s1)
// }

var (
	// size := 1024 // 这里不能用:=，简短模式只能使用在函数内部
	// max_size = size*2
	size    = 1024
	maxSize = size * 2
)

func main() {
	fmt.Println(size, maxSize)
}
