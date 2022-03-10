package main

import "fmt"

func main(){
	const (
	    a = iota
	    b = iota
	    c = iota
	)

	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c) // 2
}