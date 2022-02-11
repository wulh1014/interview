package main

func hello() []string {
	return nil
}

// 赋值
// func main() {
// 	h := hello
// 	if h == nil {
// 		fmt.Println("nil")
// 	} else {
// 		fmt.Println("not nil")
// 	}
// 	// output: not nil
// 	// 将hello函数赋值给h， h 是一个函数
// 	h1 := hello()
// 	if h1 == nil {
// 		fmt.Println("nil")
// 	} else {
// 		fmt.Println("not nil")
// 	}
// 	// output: nil
// 	// 将hello() 函数的返回值赋值给h1,h1是一个数组
// }

// 类型选择
func GetValue() int {
	return 1
}

func GetValueOfInterface() interface{} {
	return 1
}

func main() {
	// i := GetValue()
	i := GetValueOfInterface()
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
	// 编译失败，cannot type switch on non-interface value i (type int)
	// 考察类型选择 i.(type) ,i 是接口，type是固定关键字。只有接口类型能够使用类型选择
}

// init函数
// A. 一个包中，可以包含多个 init 函数；
// B. 程序编译时，先执行依赖包的 init 函数，再执行 main 包内的 init 函数；
// C. main 包中，不能有 init 函数；
// D. init 函数可以被其他函数调用；
// 答案： AB

/*
  	1、init()函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
	2、 一个包可以出现多个init函数，一个源文件也可以包含多个init函数
	3、同一个包中多个init函数的执行顺序没有明确要求，但是不同包里的init函数是根据包导入的依赖关系决定的
	4、init()函数在代码中不能被显示调用、不能被引用（赋值给函数变量），否则出现编译错误；
  	5、一个包被多个包引用，b包只会初始化一次
	6、引入包，不可出现死循坏。即 A import B,B import A，这种情况会编译失败
*/

func init() {

}

func init() {
}
