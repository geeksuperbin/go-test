package main

import "fmt"

func main() {
    //类型相同多个变量, 非全局变量
    var vname1, vname2, vname3 int8
    vname1, vname2, vname3 = 1, 2, 3
    fmt.Println(vname1,vname2,vname3) // 1 2 3

    // 和 python 很像,不需要显示声明类型，自动推断
    var vname4, vname5, vname6 = 4, 5, 6
    fmt.Println(vname4,vname5,vname6) // 4 5 6

    // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误
    vname7, vname8, vname9 := 7, 8, 9
    fmt.Println(vname7, vname8, vname9) // 7 8 9

    // 这种因式分解关键字的写法一般用于声明全局变量
    var (
        vname10 int8
        vname11 bool
    )

    fmt.Println(&vname10) // 0
    fmt.Println(vname11) // false


}

