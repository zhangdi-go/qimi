package main

import "fmt"

// 结构体的初始化
/*
注意:
	结构体定义的时候,每行后面没有逗号
	结构体初始化的时候,每行后面都有逗号
*/
type person struct {
	name string // 行尾不用写逗号
	age  int    // 行尾不用写逗号
	city string // 行尾不用写逗号
}

func main() {
	// 1. 键值对初始化
	p1 := person{
		name: "小米", // 有逗号
		age:  20,   // 有逗号
		city: "北京", // 有逗号
	}
	fmt.Printf("%#v\n", p1)

	// 2. 值的列表进行初始化
	p2 := person{
		"苹果",         // 有逗号
		50,           // 有逗号
		"California", // 有逗号
	}
	fmt.Printf("%#v\n", p2)
}
