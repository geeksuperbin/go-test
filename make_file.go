package main

import "fmt"
import "os"
import "bytes"

/*
	go 的第一个工具： 
	生成类似 000-test.go 文件
	用于小彬对 go 语言本身学习
*/
func main(){

	for i:=38; i<=99; i++ {
		// 文件名拼接
		var buffer bytes.Buffer
		buffer.WriteString(fmt.Sprintf("%d",0))
		buffer.WriteString(fmt.Sprintf("%d",i))
		buffer.WriteString("-test.go")

		//创建文件
	    file, error := os.Create(buffer.String());
	    
	    if error != nil {
	        fmt.Println(error);
	    }
	    
	    fmt.Println(file);
	    
	    file.Close();
	}
	
	
	
}