package main

import "fmt"

/*
不能给别的包的类型定义方法
类型和类型的方法必须在一个包
*/

// myInt 基于int类型自定义myInt类型
type myInt int

// SayHello 为myInt添加一个SayHello的方法
func (m myInt) SayHello() {
	fmt.Println("Hello, 我是一个myInt。")
}

func main() {
	var m1 myInt
	m1 = 100
	m1.SayHello()                 // Hello, 我是一个myInt。
	fmt.Printf("%v %T\n", m1, m1) // 100 main.myInt
}
