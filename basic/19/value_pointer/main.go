package main

import "fmt"

// 使用值接收者实现接口和使用指针接收者实现接口的区别

type animal interface {
	mover
	sayer
}

type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age  int
}

// 使用值接收者实现接口（person类型实现了mover接口）
func (p person) move() {
	fmt.Printf("%s在跑...\n", p.name)
}

// 使用指针接收者实现接口（person类型实现了sayer接口）
func (p *person) say() {
	fmt.Printf("%s在说...\n", p.name)
}

func main() {
	// 定义一个接口变量a
	var a animal

	p1 := person{ // p1是person类型的值变量
		name: "小草莓",
		age:  20,
	}
	a = &p1
	a.move()
	a.say()
	fmt.Printf("a: %v\n", a)

	p2 := &person{ // p2是person类型的指针变量
		name: "香蕉君",
		age:  30,
	}
	a = p2
	a.move()
	a.say()
	fmt.Printf("a: %v\n", a)
}
