package main

import "fmt"


func main(){
    
    // 数组
    var arr [6]string = [6]string{"aiha","ban","geekbin","do","gogo","best"}
    //   0      1       2       3    4      5
    // "aiha","ban","geekbin","do","gogo","best"
    
    //  切片的定义 slice
    // var brr []string


    p := fmt.Println

    p("[1:3]", arr[1:3])
    p("[:3]", arr[:3])
    p("[1:]", arr[1:])
    p("[1:3]", arr[1:3])
    p("[:]", arr[:])

 	// [1:3] [ban geekbin]
	// [:3] [aiha ban geekbin]
	// [1:] [ban geekbin do gogo best]
	// [1:3] [ban geekbin]
	// [:] [aiha ban geekbin do gogo best]


    //   0      1       2       3    4      5
    // "aiha","ban","geekbin","do","gogo","best"	
	s1 := arr[1:3]

	// 切片后重排索引
	//   0      1        2    3      4     5
	// "ban","geekbin","do","gogo","best"  ?	
	s2 := s1[2:5]  // s2 对 s1 的再次切片

	p(s1) // [ban geekbin]
	p(s2) // [do gogo best]
	
	// s2  是切片的切片	
    //      0   1    2
	// s2: do gogo best
	


	// 切片的长度和容量
	p("s1的长度：", len( s1)) // 2
	p("s1的容量：", cap( s1)) // 5
	p("s1的长度：", len( s2)) // 3
	p("s1的容量：", cap( s2)) // 3


	// 切片是对数组的局部或全部的一种引用形式
	
	// 切片的操作
	// 1) 切片的拓展追加 append
	
	d := [...]int{111,222,333,444,555,666}
	//  0    1  2   3    4  5
	// 111,222,333,444,555,666
	
	d1 := d[:3] 
	p(d1)
	//  0   1   2
	// 111 222 333


	// d1 = append(d1, 777)

	// p(d1) [111 222 333 777]
	
	d1 = append(d1, 777)
	p(d1) // [111 222 333 777]
	p(d) // [111 222 333 777 555 666] 将索引 3 的 444 覆写成 777


	// d2 := d[:3]
	// p(d2) // 111 222 333
	// d2 = append(d2, []int{777, 888, 999}...) // 切片展开
	// p(d2)





    
}