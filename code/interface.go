package main

import "fmt"

type Animals interface {
	Run() string
	Call() string
}

type Cats struct {
	Name string
}

func (c *Cats) Call() string{
	return c.Name
}

func (c *Cats) Run() string{
	return "Run"
}

func main() {
	c := &Cats{"cat"}
	var a Animals
	a = c
	fmt.Println(a.Call())
}
