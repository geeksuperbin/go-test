package main

import "fmt"


func main(){
	a, b := swap("Google", "Baidu")
	fmt.Println(a,b) // Baidu Google
}

// 函数返回多个值
func swap(x, y string) (string, string) {
	return y, x
}

