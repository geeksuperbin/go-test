package main

import "fmt"


func main(){
	// 循环嵌套输出 2 到 100 间的素数即质数
	for i:=2; i<=100; i++ {
		for j:=2; j<=i; j++ {
			if(i%j == 0){
				if(i==j){ // 判断是否是最后一位，如果是就直接输出
					fmt.Println(i)
				}else{
					break
				}
			}
		}
	}
}