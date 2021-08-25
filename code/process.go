package main

import "fmt"

type Student struct {
	Name 	string
	Age 	int
}

func main()  {
	m := make(map[string]*Student)

	l := []Student{
		{"aaa", 20},
		{"bbb", 30},
		{"ccc", 40},
	}

	for _, stu := range l {
		m[stu.Name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", v)
	}
	//continueDemo()
	gotoDemo2()
}

func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}

func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}