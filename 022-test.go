package main

import "fmt"


func main(){

	// 方法1： 计算 1到 10 的数字之和
	sum := 0
	for i:=1; i<=10; i++ {
		sum += i
	}

	fmt.Printf("1-10的求和的值为: %d \n",sum) // 1-10的求和的值为: 55 


	// 方法2：计算 1到 10 的数字之和
	sum2 := 0
	i2 := 0
	for i2 <= 10 {
		sum2 += i2
		i2++
	}

	fmt.Println(sum2) // 55

}