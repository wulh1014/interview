package main

// // 下面这段代码能否通过编译，不能的话原因是什么；如果能，输出什么。

// func main() {
//
// 	list := new([]int)
// 	list = append(list, 1)
// 	fmt.Println(list)
// 	// 不能通过编译，new([]int) 之后的list是一个 *[]int类型的指针，而append的第一个参数需要接收一个slice类型 (first argument to append must be slice; have *[]int)
// 	// 可以使用make初始化之后在进行使用，或 list := []int 进行使用
// }

// // 下面这段代码能否通过编译，如果可以，输出什么？
// func main() {
//
// 	s1 := []int{1, 2, 3}
// 	s2 := []int{4, 5}
// 	s1 = append(s1, s2)
// 	fmt.Println(s1)
//
// 	// 不能通过编译，append的第二个参数需要是type int，而不能是slice (cannot use s2 (type []int) as type int in append)
// 	// 可以使用 ...操作符将一个切片追加到另一个切片上;append(s1,s2...) 或者，将元素一个个进行追加 append(s1,1,2,3)
// }

// // 下面这段代码能否通过编译，如果可以，输出什么？
// var (
// 	size := 1024
// 	max_size = size * 2
//
// 	// 不能通过编译，go语言支持简短声明 如 x := 100,但该声明方式有限制
// 	// 必须使用显示初始化  (必须得给变量赋值，而不是由编译器隐式的给予变量类型的默认值)
// 	// 不需要提供数据类型，编译器会自动推导
// 	// 只能在函数内部使用简短声明
// )
//
// func main() {
// 	fmt.Println(size, max_size)
//
// }
