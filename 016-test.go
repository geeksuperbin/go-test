package main

import "fmt"

func main(){
	// 进制转换工具：https://tool.lu/hexconvert/

	a := 60 // 00111100
	b := 13 // 00001101
	fmt.Println(a&b) // 十进制表示：12 二进制表示 00001100

	fmt.Println(a|b) // 61  00111101

	fmt.Println(a^b) // 49  00110001u
}