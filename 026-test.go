package main

import "fmt"


func main(){
	for i:=1; i<=10; i++ {
		if(i == 5){
			goto LABEL
		}else{
			fmt.Println(i)
		}
	}

	LABEL:
		fmt.Println("goto")

	/*
    输出结果：
			1
			2
			3
			4
			goto
	*/
}