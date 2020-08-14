package main

import "fmt"

func hello() []string {
	return nil
}

func main() {
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	k := hello()
	if k == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

// func GetValue() int {
// 	return 1
// }

// func main() {
// 	i := GetValue()
// 	switch i.(type) {
// 	case int:
// 		println("int")
// 	case string:
// 		println("string")
// 	case interface{}:
// 		println("interface")
// 	default:
// 		println("unknown")
// 	}
// }
