package main

import (
	"fmt"
)

func main() {
	s := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	p := &s

	fmt.Println(*p)
	fmt.Println(p.name)  // 取到name的值
	fmt.Println(&p.name) // 取到name的地址
	fmt.Println((*p).name)
	// fmt.Println(p->name) // 语法错误
}
