/*
	package main 用于定义包名
	必须在源文件中非注释的第一行指明这个文件属于哪个包
	package main 表示是一个可独立执行的程序
	每个 Go 应用都包含一个名为 main 的包

*/
package main

/*
	import "fmt",用于告诉 Go 编译器，这个程序需要使用 fmt 包中的函数
	fmt 包实现了格式化 IO 输入输出 的函数
*/
import "fmt"


func init(){
	fmt.Println("go go go")
}

/*
	func main() 是程序开始执行的函数
	main 函数是每一个可执行程序所必须包含的
	是在 init 后执行的函数
*/
func main(){

	/*
		fmt.Println() 可以将字符串输出到控制台，并在最后
		自动增加换行符 \n
	*/
	fmt.Println("Hello World!!");
}