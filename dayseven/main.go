package main

// 1、关于字符串连接，下面正确的是:
// A. str := ‘abc’ + ‘123’
// B. str := “abc” + “123”
// str := ‘123’ + “abc”
// D. fmt.Sprintf(“abc%d”, 123)

// func main() {
// 	// str := 'abc' + '123'
// 	// '' 括起来为rune类型，不是字符串类型
//
// 	str := "abc" + "123"
// 	str1 := fmt.Sprintf("abc%d", 123)
// 	str2 := strings.Join([]string{"11", "22"}, ",")
//
// 	buf := bytes.NewBufferString("123")
// 	buf.WriteString("abc")
// 	fmt.Println(str, str1, str2, buf.String())
//
// 	// 考察字符串拼接方法， + , fmt.sprintf() , strings.Join(), buffer.WriteString()
// }

// 2、 下面代码能否编译通过？如果可以，输出什么？

// const (
// 	x = iota
// 	_
// 	y
// 	z = "zz"
// 	k
// 	p = iota
// )
//
// func main() {
// 	fmt.Println(x, y, z, k, p)
//
// 	// 编译通过，输出 0, 2 zz zz 5 考察iota的使用
// 	// https://www.cnblogs.com/zsy/p/5370052.html
// }

// 3、 下列赋值正确的是

func main() {
	// var x = nil   // false
	// var x interface{} = nil   // true
	// var x string = nil  // false
	// var x error = nil   // true

	// nil 值只能赋值给指针、chan， map， slice， func ，interface 等指针类型变量
	// error 类型是一个内置的接口
	// The error built-in interface type is the conventional interface for
	// // representing an error condition, with the nil value representing no error.
	// type error interface {
	// 	Error() string
	// }

}
