package main

import (
	"fmt"
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

func (e empty) Say() {
	fmt.Println("empty")
}
//var a struct{}
//var b empty
//var c *empty

//func main() {
//	fmt.Println(unsafe.Sizeof(a))
//
//	s := &empty{}
//	s.Say()
//	fmt.Println(unsafe.Sizeof(s))
//
//	b = empty{}
//	b.Say()
//	fmt.Println(unsafe.Sizeof(b))
//
//	c = &empty{}
//	b.Say()
//	fmt.Println(unsafe.Sizeof(c))
//
//}


