package main

import "fmt"

// 下面这段代码能否通过编译？不能的话，原因是什么？如果通过，输出什么？
func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		name string
		age  int
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}

	/*
		不能编译通过   invalid operation: sm1 == sm2
		该题目考察结构体的比较:
		1、 结构体只能比较是否相等，但是不能比较大小
		2、相同类型的结构体才能够进行比较，结构体是否相同不但与属性类型有关，还与属性顺序有关
		3、 如果struct的所有成员都可以比较，则该struct就可以通过== 或!=进行比较是否相等，比较时逐个项进行比较，如果每一项都相等，则两个结构体才相等，否则不相等
		4、bool、数值类型、字符、指针、数组等是可以比较的
		5、切片、map、函数等是不能比较的。
		6、参考文档 https://golang.org/ref/spec#Comparison_operators

	*/
}
