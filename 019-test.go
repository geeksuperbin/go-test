package main

import "fmt"

func main(){
	// if 嵌套语句，可在 if else if 语句进行嵌套
	var a int = 10

	if(a > 20){
		fmt.Println("a 大于 20")
	}else if(a >= 10){
		fmt.Println("a 大于等于 10")
	}else{
		fmt.Println("a 小于 10")
	}
}