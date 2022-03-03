package main;

import "fmt";

func main(){
    //  %s 字符串格式，%d 十进制的整数格式 %f 浮点数
    var res = fmt.Sprintf("Hello,World-%s-%d-%f","geekbin",2022,666.0);
    fmt.Println(res);
}   
