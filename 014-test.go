package main

import "fmt"


func main(){
	var a int = 10
	var b int = 20

	// == a,b 不相等
	if(a == b){
		fmt.Println("a,b 相等")
	}else{
		fmt.Println("a,b 不相等")
	}

	// != a,b 不相等
	if(a != b){
		fmt.Println("a,b 不相等")
	}


	// > a小于b
	if(a > b){
		fmt.Println("a大于b")
	}else{
		fmt.Println("a小于b")

	}

	// < a小于b
	if(a < b){
		fmt.Println("a小于b")
	}

	// >=  a小于等于b
	if(a >= b){
		fmt.Println("a大于等于b")
	}else{
		fmt.Println("a小于等于b")

	}

	// <= a小于等于b
	if(a <= b){
		fmt.Println("a小于等于b")
	}

}