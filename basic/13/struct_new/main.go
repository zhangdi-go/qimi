package main

import "fmt"

// 结构体的构造函数: 构造一个结构体实例的函数
// 结构体是值类型

type person struct {
	name string
	age  int
	city string
}

// 构造函数
func newPerson(name string, age int, city string) *person {
	return &person{
		name: name, // 前面的name是结构体的name字段, 后面的name是传入函数的参数
		age:  age,
		city: city,
	}
}

func main() {
	p1 := newPerson("小米", 20, "北京")
	fmt.Printf("Type:%T\nValue:%#v\n", p1, p1)
}
