package main

// 1、通过指针p 访问其变量的方式
// A.p.name
// B.(&p).name
// C.(*p).name
// D.p->name

// func main() {
// 	var p = &struct {
// 		name string
// 	}{}
//
// 	fmt.Println(p.name)
// 	fmt.Println((*p).name)
//
// 	// & 取值运算符
// 	// * 指针解引用
//
// }

// 2、类型别名与类型定义
// 下面代码是否可以编译通过
type MyInt1 int   // 类型定义
type MyInt2 = int // 类型别名

func main() {
	// 	var i int = 0
	// 	var i1 MyInt1 = i
	// 	var i2 MyInt2 = i
	// 	fmt.Println(i1, i2)

	// 不能编译通过cannot use i (type int) as type MyInt1 in assignment
	// i 是int类型，i1是MyInt1类型  i2 的类型是MyInt2，但是MyInt2 只是int类型的别名，本质上还是int类型
	// go是强类型语言，不同类型之间赋值需要进行类型转换
	// var i1 MyInt1 = MyInt(i)
}
