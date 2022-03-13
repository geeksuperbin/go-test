package main

import "fmt"


func main(){
	
	fmt.Println(max(5,10)) // 10
	
}


	// max() 函数，该函数传入两个整型参数 num1 和 num2 
	// 并返回这两个参数的最大值

	func max(num1, num2 int) int {
		// 声明局部变量
		var result int
		if(num1 > num2){
			result = num1
		}else{
			result = num2
		}
		return result
	}