package main

import "fmt"

// 结构体的继承

// Animal 结构体
type Animal struct {
	name string
}

// Animal 的方法
func (a *Animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

// Dog 结构体
type Dog struct {
	Feet    int
	*Animal // 匿名嵌套，而且嵌套的是结构体指针
}

// Dog 的方法
func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "大黄",
		},
	}
	d1.wang() // 狗会wang()
	d1.move() // 狗会move()
}
