package main

import "fmt"

// 闭包

// 定义一个函数，它的返回值是函数
// 把函数作为返回值
func a() func() {
	name := "a()的内部变量"
	return func() {
		fmt.Println("hello", name)
	}
}

func main() {
	r := a() // a函数的返回值赋值给了r
	r()      // 相当于执行了a函数内部的匿名函数
}
