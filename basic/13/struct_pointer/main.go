package main

import "fmt"

// 结构体指针

type person struct {
	name string // 行尾不用写逗号
	age  int    // 行尾不用写逗号
	city string // 行尾不用写逗号
}

func main() {
	var p1 = new(person)
	fmt.Printf("Type p1: %T\n", p1)
	(*p1).name = "小草莓" // 标准写法
	p1.age = 26        // 语法糖
	p1.city = "上海"     // 语法糖
	fmt.Printf("p1: %#v\n", p1)

	// 取结构体的地址进行实例化
	p2 := &person{}
	fmt.Printf("Type p2: %T\n", p2)
	p2.name = "las fresas"
	p2.age = 30
	p2.city = "Auckland"
	fmt.Printf("p2: %#v\n", p2)
}
