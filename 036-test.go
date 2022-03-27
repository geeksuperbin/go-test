package main

import "fmt"

func main() {

	// 数组
	var arr [6]string = [6]string{"aiha", "ban", "geekbin", "do", "gogo", "best"}
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
	s2 := s1[2:5] // s2 对 s1 的再次切片

	p(s1) // [ban geekbin]
	p(s2) // [do gogo best]

	// s2  是切片的切片
	//      0   1    2
	// s2: do gogo best

	// 切片的长度和容量
	p("s1的长度：", len(s1)) // 2
	p("s1的容量：", cap(s1)) // 5
	p("s1的长度：", len(s2)) // 3
	p("s1的容量：", cap(s2)) // 3

	// 切片是对数组的局部或全部的一种引用形式

	// 切片的操作
	// 1) 切片的拓展追加 append

	d := [...]int{111, 222, 333, 444, 555, 666}
	//  0    1  2   3    4  5
	// 111,222,333,444,555,666

	d1 := d[:3]
	p(d1)
	//  0   1   2
	// 111 222 333

	// d1 = append(d1, 777)

	// p(d1) [111 222 333 777]

	// d1 = append(d1, 777)
	// p(d1) // [111 222 333 777]
	// p(d) // [111 222 333 777 555 666] 将索引 3 的 444 覆写成 777

	d2 := d[:3]
	// p(d2) // 111 222 333
	d2 = append(d2, []int{777, 888, 999}...) // 切片展开
	p(d2)
	d2 = append(d2, []int{44, 55, 66}...) // 切片展开
	p(d2)                                 // d2 长度增加 新增了一个数组，原来数组 d 由于容量大小，则不再使用 [111 222 333 777 888 999 44 55 66]
	p(d)                                  // [111 222 333 777 888 999]

	p("===========================")

	// 切片的复制  copy() 函数
	//             0    1   2  3   4   5
	f := [...]int{111, 222, 333, 444, 555, 666}
	p(f)
	f1 := f[1:3] // 222  333
	p("f1", f1)
	f2 := f[3:] //  444  555  666
	p("f2", f2)

	copy(f2, f1)            // 将 f1 切片 拷贝到 f2 即 将 f1 的前两个值覆盖给 f2
	p("将 f1 切片 拷贝到 f2", f2) // 222  3333  666

	p(f) // [111 222 333 222 333 666], 能看出来切片的操作，会对原数组产生音响，切片是对数组的引用

	t := make([]int, 4, 5)
	t2 := make([]int, 4)
	p(t)
	p(len(t), cap(t)) // 4 5
	p(t2)
	p(len(t2), cap(t2)) // 4 4

}
