package main

import "fmt"

func hello(i int) {
	fmt.Println(i)
}
func main() {
	i := 5
	defer hello(i)
	i = i + 10
}

// type People struct{}

// func (p *People) ShowA() {
// 	fmt.Println("showA")
// 	p.ShowB()
// }
// func (p *People) ShowB() {
// 	fmt.Println("showB")
// }

// type Teacher struct {
// 	People
// }

// func (t *Teacher) ShowB() {
// 	fmt.Println("teacher showB")
// }

// // 观察结构体嵌套
// func main() {
// 	t := Teacher{}
// 	t.ShowA()
// 	t.ShowB()
// 	t.People.ShowA()
// 	t.People.ShowB()
// }
