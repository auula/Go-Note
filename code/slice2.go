package main

import "fmt"

func hello(num []int) {
	num[0] = 18
}

func main() {
	i := []int{5, 6, 7}
	hello(i)
	fmt.Println(i[0])

}

func Test() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	a = append(a, b)
	fmt.Println(a)
}

