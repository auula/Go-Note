package main

import "fmt"

func main() {
	//SliceTest()
	s1 := make([]int,1024)
	s1 = append(s1,111,112)
	fmt.Println(cap(s1))
}


func SliceTest() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}