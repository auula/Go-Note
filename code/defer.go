package main

import "fmt"

func AA()  {
	defer func() {
		fmt.Println(1)
	}()
	fmt.Println("aa")
}

func BB()  {
	defer func() {
		fmt.Println(2)
	}()
	fmt.Println("bb")
}

func CC()  {
	defer func() {
		fmt.Println(3)
	}()
	fmt.Println("cc")
}

//func main() {
//	fmt.Println("start")
//	AA()
//	BB()
//	CC()
//	fmt.Println("end")
//}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//func main() {
//	x := 1
//	y := 2
//	defer calc("AA", x, calc("A", x, y))
//	x = 10
//	defer calc("BB", x, calc("B", x, y))
//	y = 20
//}

