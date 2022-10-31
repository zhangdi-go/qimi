package main

import "fmt"

// 函数之变量作用域

// 定义全局变量num
var num int = 1

// 定义函数
func testGlobal() {
	num := "小草莓坏笑"
	// 可以在函数中访问全局变量
	// 1. 先在自己函数中查找局部变量
	// 2. 找不到就往外找全局变量
	fmt.Printf("变量num: %v\n", num)
}

func main() {
	testGlobal()
}
