package main

import "fmt"


func main(){

	// switch 默认情况下 case 最后自带 break 语句
	var marks int = 70

	switch marks {
	case 90:
		fmt.Println("优秀")
	case 80:
		fmt.Println("良好")	
	case 70:
		fmt.Println("及格")
		fallthrough
	default:
		fmt.Println("请家长吧")	
	}
}

/*
	输出结果：
	
	及格
	请家长吧

*/