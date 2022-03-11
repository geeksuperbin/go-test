package main

import "fmt"

func main(){

	// 返回变量存储地址
	var a int = 2
	fmt.Println(&a) // 0xc0000aa058

	// * 指针变量
	var b *int
	b = &a
	fmt.Println(b) // 0xc0000aa058

}