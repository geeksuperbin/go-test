package main

import "fmt"

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	// 面积大小
	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d",area) // 面积为：50
}