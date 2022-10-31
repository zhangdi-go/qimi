package main

import "fmt"

// 方法是有接收者的特殊函数

// Person 是一个结构体
type Person struct {
	name string
	age  int
}

// NewPerson 是一个Person类型的构造函数，返回指针类型的Person
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name, // 前面的name是结构体的name字段,后面的name是传入函数的参数
		age:  age,  // 前面的age是结构体的age字段,后面的age是传入函数的参数
	}
}

// 1. 值接收者

// Dream 是Person类型的方法。
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言\n", p.name)
}

// SetAge1 是Person类型的修改年龄的方法。
func (p Person) SetAge1(newAge int) {
	p.age = newAge
}

// 2. 指针接收者

// SetAge2 是Person类型的修改年龄的方法。
func (p *Person) SetAge2(newAge int) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("小草莓", 20)
	(*p1).Dream()
	p1.Dream()
	fmt.Printf("p1.age: %v\n", p1.age)
	p1.SetAge1(1000)
	fmt.Printf("p1.age: %v\n", p1.age)
	p1.SetAge2(9000)
	fmt.Printf("p1.age: %v\n", p1.age)
}
