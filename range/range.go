package main

import "fmt"

func main() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	i := 1
	for _, val := range slice {
		m[i] = &val
		i++
		/*
		错误写法
		output:
		3 -> 3
		0 -> 3
		1 -> 3
		2 -> 3
		参考原因：
		在编译期将原切片或者数组赋值给一个新变量 ha
		for range 循环的时候会创建每个元素的副本，而不是元素的引用，所以&val取得都是变量val的地址，所以最后的到的值是3，所以输出都是3
		*/
		// // 正确写法
		// // m[key] = &a[index]
		// value := val
		// m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}

}
