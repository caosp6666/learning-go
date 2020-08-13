package main

import "fmt"

// func funcMui(x, y int) (sum int, error) {
// 	return x + y, nil
// }
// 当返回值有两个的时候，不能一个命名一个不命名，会报错

// 1.
func main() {
	s := make([]int, 5) // 生成的是5个0的切片
	fmt.Println(s)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	// a, b := funcMui(1, 4)
	// fmt.Println(a, b)
}

// //2.
// func main() {
// 	s := make([]int, 0)
// 	s = append(s, 1, 2, 3, 4)
// 	fmt.Println(s)
// }
