package main

import "fmt"

// func main() {
// 	a := 5
// 	b := 8.1
// 	// fmt.Println(a + b) // 不同类型不能相加
// 	fmt.Println(float64(a) + b)
// }

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0])
}

// func main() {
// 	a := [2]int{5, 6}
// 	b := [3]int{5, 6}
// 	if a == b { // a b类型不同，不能比较
// 		fmt.Println("equal")
// 	} else {
// 		fmt.Println("not equal")
// 	}
// }
