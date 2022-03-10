package main

import "fmt"

func main(){
	// var a int = 1
	// var b int = 2

	// 相加
	fmt.Println(1 + 2) // 3
	// 相减
	fmt.Println(0.3 - 0.1) //i0.2
	// 相乘
	fmt.Println(0.3 * 0.4) // 0.12
	// 相除
	fmt.Println(0.6 / 0.2) // 3
	// 求余
	fmt.Println(10 % 3) // 1
	// 自增
	var i int = 9
	i++
	fmt.Println(i) // 10 , i++ 不能直接放到 fmt.Println()括号中会报错
	var j int = 12
	j--
	fmt.Println(j) // 11

}