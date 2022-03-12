package main

import "fmt"


func main(){
	
	for i:=1; i<=10; i++ {
		if(i == 6){
			continue
		}else{
			fmt.Println(i)
		}
	}

	/*
	 输出结果：
	    1
		2
		3
		4
		5
		7
		8
		9
		10

	*/
	
}