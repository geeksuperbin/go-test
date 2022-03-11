package main

import "fmt"


func main(){
	a := 1
	b := 2

	// = 简单的赋值运算符，将一个表达式的值赋值给一个左值
	c := a + b
	fmt.Println(c) // 3

	// += 相加后再赋值
	var d int
	d += a
	fmt.Println(d) // 1

	// -= 相减后再赋值
	var e int
	e -= b
	fmt.Println(e) // -2

	// *= 相乘后再赋值
	b *= e
	fmt.Println(b) // -4

	// /= 相除后再赋值
	b /= a
	fmt.Println(b) // -4

	// %= 求余后再赋值
	var f int = 10
	f %= 3
	fmt.Println(f) // 1

	// <<= 左移后赋值
	c <<= 2
	fmt.Println(c) // 12

	// >>= 右移后赋值
	c >>= 2
	fmt.Println(c) // 3

	// &= 按位与后赋值
	c &= 2
	fmt.Println(c) // 2

 	// ^= 按位异或后赋值
 	c ^= 2
 	fmt.Println(c) // 0

 	// |= 按位或后赋值
 	c |= 2
 	fmt.Println(c) // 2

}