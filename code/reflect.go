package main

import (
	"fmt"
	"reflect"
)


func main() {
	var i int64 = 100
	v := reflect.ValueOf(&i)
	//v1 := reflect.ValueOf(i)
	v.Elem().SetInt(10)
	//v1.Elem().SetInt(11)
	fmt.Println(i)
}