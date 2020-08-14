package main

import "fmt"

// func main() {
// 	var i interface {
// 		Error() error
// 		Sayhi() string
// 	}
// 	if i == nil {
// 		fmt.Println("nil")
// 	} else {
// 		fmt.Println("not nil")
// 	}

// 	fmt.Println("value of i is:", i)
// 	fmt.Printf("type of i is %T\n", i)

// 	type ES interface {
// 		Error() error
// 		Sayhi() string
// 	}
// 	var es ES
// 	fmt.Printf("type of es is %T\n", es)
// }

// 当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil。

func main() {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
}

// 删除map中不存在的键值对不会报错，但是也不会有任何效果
