package main

import "fmt"

// 1、关于channel的说法正确的是
// A. var ch chan int
//
// B. ch := make(chan int)
//
// C. <- ch
//
// D. ch <-
// channel 的声明， 读写channel
// 写channel时必须要有值
// func main() {
// 	var ch chan int
// 	ch1 := make(chan int)
// 	<-ch
// 	ch1 <- 5
// }

// 2、下面代码输出什么？
// type person struct {
// 	name string
// }
//
// func main() {
// 	var m map[person]int
// 	p := person{"mike"}
// 	fmt.Println(m[p])
// 	// output: 0
// 	// 打印一个map中不存在的值，返回元素类型的零值
// }

// 3、下面代码输出什么

func hello(num ...int) {
	num[0] = 18
}

func main() {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i)
	// i: [18,6,7]
	fmt.Println(i[0])

	i1 := 1
	hello(i1)
	fmt.Println(i1)
	// output: i1 =1
	// 可变参数函数
}
