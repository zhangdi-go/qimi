package main

import "fmt"

// 闭包

// 定义一个函数，它的返回值是函数
func a(name string) func() {
	return func() {
		fmt.Println("hello", name)
	}
}

func main() {
	// 闭包 = 函数 + 外部变量的引用
	r := a("外部变量的引用") // r此时就是一个闭包
	r()               // 相当于执行了a函数内部的匿名函数
}
