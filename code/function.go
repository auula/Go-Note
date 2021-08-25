package main

import (
	"fmt"
)


func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

type fn func(int, int) int

//func main() {
//	sum(2,2,add, sub)
//}

func sum (x, y int, fns ...fn) {
	for _, v := range fns {
		fmt.Println(v(x, y))
	}
}

type Job func()


func main() {
	ch := make(chan Job,3)
	defer close(ch)
	go func() {
		for {
			fn, ok := <- ch
			if !ok {
				return
			}
			fn()
		}
	}()

	var job1 = func() {
		fmt.Println(1)
	}
	var job2 = func() {
		fmt.Println(2)
	}
	var job3 = func() {
		fmt.Println(3)
	}
	var job4 = func() {
		fmt.Println(4)
	}
	ch<- job1
	ch<- job2
	ch<- job3
	ch<- job4
	fmt.Println("success")
}
