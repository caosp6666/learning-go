package main

import "fmt"

func main() {
	i := -5
	j := +5
	fmt.Printf("%d %d\n", i, j)
	fmt.Printf("%+d %+d\n", i, j)
	// 注意对比区别，+代表输出符号，负号自然情况下也会输出
}
