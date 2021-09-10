package main

import "fmt"


//断言

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

type Dogs struct {

}

func (d *Dogs) Call() string  {
	return "dogs call"
}

func (d *Dogs) Run() string {
	return "dog run"
}

func main() {
	z := &Cats{"cat"}
	var b Animals
	b = z
	if v, ok := b.(*Dogs); ok {
		fmt.Println(v)
	}else{
		fmt.Println("error")
	}
}
