package main

import "fmt"

func main() {
  a := 1
  b := 2
  var c int
  var d int
  // 并行，同时赋值
  // 要交换两个变量的值，则可以简单地使用 a, b = b, a，两个变量的类型必须是相同
  c,d = a,b
  fmt.Println(c)
  fmt.Println(d)
}