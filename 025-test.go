package main

import "fmt"


func main(){
	
	for i:=1; i<=10; i++ {
		if(i == 5){
			break
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
	*/
}