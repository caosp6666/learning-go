package main

import "fmt"

var ii int
var jj = 20

// kk := "haha"  // 这样定义是错误的 "non-declaration statement outside function body"

func main() {
	// 1.
	var i int // int 默认为0

	// 2.
	var j = 20

	// 3.
	k := "haha"

	fmt.Println(i, ii, j, jj, k)
}
