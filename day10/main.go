package main

// 1、不同类型操作
// func main(){
// 	a:=5
// 	b := 8.1
// 	fmt.Println(a + b)
// 	// a 是 int， b 是float  不同类型之间的数值不能相加
//  output: compilation error
// }

// 2、下面这段代码输出什么
// func main() {
// 	a := [5]int{1, 2, 3, 4, 5}
// 	t := a[3:4:5]
// 	t2 := a[4:5]
// 	fmt.Println(t[0],t2)
// 	// output: 4
// 	// output: t2 = 5
// 	t1 := a[3:4]
// 	fmt.Println(cap(t1))
// 	// output: 2
//
// 	// 考察点： 操作符[i:j]
// 	// 基于数组或者切片可以使用操作符[i:j]创建新切片，从索引i开始，到索引 j 结束
// 	// 新切片的值包含原数组切片索引 i 的值
// 	// 但不包含索引j的值
// 	// i,j 是可选的，i默认是0，j默认是数组的长度
// 	// 假如底层数组的大小为k，截取之后获得的切片的长度和容量，长度： j-i 容量: k - i
// 	// [i:j]截取操作符还可以有第三个参数，形如[i:j:k]，k用来限制原切片的容量，但不能超过原切片的底层数组大小。新切片的长度和容量分别是：j-i、k-i
// 	// 所以切片t的 长度和容量都为1，t1的长度和容量为1，2
// }

// 3、下面这段代码输出什么？
// func main() {
// 	a := [2]int{5, 6}
// 	b := [3]int{5, 6}
// 	if a == b {
// 		fmt.Println("equal")
// 	} else {
// 		fmt.Println("not equal")
// 	}
// 	//  invalid operation: a == b (mismatched types [2]int and [3]int)
// 	// go中的数组是值类型，可以比较
// 	// 数组的长度和元素的类型也是数组类型的一部分，a 和 b的数组类型不同，所以不是同一个类型，不能比较
// }
