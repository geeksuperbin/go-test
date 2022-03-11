package main

import "fmt"

func main(){
	var a bool = true
	var b bool = false

	// 逻辑 AND 运算符
	fmt.Println(a && b) // false，go 也有熔断机制
	// 逻辑 OR 运算符
	fmt.Println(a || b) // true
	//逻辑 NOT 运算符
	fmt.Println(!a) // false
	fmt.Println(!b) // true
}