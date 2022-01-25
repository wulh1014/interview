package main

func main() {
	// // 1、下列代码，输出什么
	// s := make([]int, 5)
	// fmt.Println(s, cap(s), len(s))
	// // cap : 5 len : 5
	// s = append(s, 1, 2, 3)
	// fmt.Println(s, cap(s), len(s))
	// // cap: 10, len(8)
	// // output
	// // [0 0 0 0 0 1 2 3]
	// // 原因：make函数初始化了长度为5的slice，其中每一个元素都是其类型的默认值，因此append追加元素后会输出(0 0 0 0 0 1 2 3)
	// s2 := make([]int, 0)
	//
	// s2 = append(s2, 1, 2, 3, 4)
	// fmt.Println(s2)
	// // output [1 2 3 4]
	// // 原因: s2初始化了一个没有元素的空数组[],因此在append追加元素时从下标0开始追加

	// 2、下列代码有什么问题
	// func funcMui(x,y int) (sum int,error){
	// 	return x+y,nil
	// }
	// 原因：第二个返回者值error没有命名
	// go语言中可以不给返回值命名，但是多个返回值中如果有一个给返回值命名，其他的也要命名

	// 3、new 和 make 的区别
}
