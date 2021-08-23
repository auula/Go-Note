package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("a:%d, prt : %p\n", a, &a)	//a:10, prt : 0xc00000e0a8
	fmt.Printf("b:%p, type:%T\n", b, b)		//b:0xc00000e0a8, type:*int
	fmt.Println(b)		//0xc00000e0a8
	fmt.Println(&b)		//0xc000006028
	fmt.Println(*b)		//10
}