package main

import "fmt"

// 为什么需要接口？

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵！")
}

type dog struct {
	name string
}

func (d dog) say() {
	fmt.Println("汪汪汪！")
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Println("fuck!")
}

/*
在Go语言中接口（interface）是一种类型，一种抽象的类型。
接口不管是什么类型, 它只管要实现什么方法
*/

// sayer 只要实现了say()这个方法的类型就可以成为sayer类型
type sayer interface {
	say()
}

// da 这个函数需要接收一个通用的类型，并且这个类型有个say方法
func da(arg sayer) {
	arg.say() // 不管传进来的是什么，都要打。 打Ta, Ta就会叫，就要执行Ta的say方法
}

func main() {
	c1 := cat{}
	da(c1)
	d1 := dog{}
	da(d1)
	p1 := person{}
	da(p1)

	fmt.Println()

	// s是sayer类型的接口
	var s sayer
	c2 := cat{name: "小黄"}
	s = c2 // s可以接收cat类型
	fmt.Println(s)
	da(s)
	p2 := person{name: "Elles"}
	s = p2 // s也可以接收person类型
	fmt.Println(s)
	da(s)
}
