package main

import "fmt"

// func main() {
// 	i := -5
// 	j := +5
// 	fmt.Printf("%d %d\n", i, j)
// 	fmt.Printf("%+d %+d\n", i, j)
// 	// 注意对比区别，+代表输出符号，负号自然情况下也会输出
// }

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

// 观察结构体嵌套
func main() {
	t := Teacher{}
	t.ShowB()
	t.People.ShowA()
	t.People.ShowB()
}
