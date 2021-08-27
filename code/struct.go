package main

import (
	"fmt"
	"unsafe"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

//func main() {
//	t := Teacher{}
//	t.ShowA()
//}

type Animal struct {
	Name string
}

func (a *Animal) Move()  {
	fmt.Printf("%s move\n", a.Name)
}

type Cat struct {
	Name 	string
	Animal
}

func (c *Cat) Call() {
	fmt.Printf("%s call\n", c.Name)
}

//func main() {
//	c := &Cat{
//		Name:   "cat",
//		Animal: Animal{"animal"},
//	}
//	c.Move()
//	c.Call()
//	fmt.Println(c.Name)
//	fmt.Println(c.Animal.Name)
//}


type empty struct{}


var a struct{}
var b empty
var c *empty

func main() {
	fmt.Println(unsafe.Sizeof(a))

	s := &empty{}
	fmt.Println(unsafe.Sizeof(s))

	b = empty{}
	fmt.Println(unsafe.Sizeof(b))

	c = &empty{}
	fmt.Println(unsafe.Sizeof(c))

}


type T struct {
	a int8		//1
	b int32		//4
	c int16		//2
	//d int64		//8

}

//func main() {
//	t := T{}
//	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(t), unsafe.Alignof(t))
//
//}
/**a 是占用1字节，b占用4字节，c占用2字节，这个结构体里，最大的是4字节，以最大字节作为读取边界，所以每次是读4字节，a占用1字节，b占用4字节，无法一次读取，所以就是|axxx|bbbb}
c占2字节 所以当前内存是：|axxx|bbbb|ccxx| 12%4 = 0，所以这个t占用12字节
**/
