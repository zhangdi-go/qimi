package main

import "fmt"

// 函数可以作为变量

// 定义全局变量num
var num int = 10

// 定义函数
func testGlobal() {
	num := "小草莓坏笑"
	// 可以在函数中访问全局变量
	// 1. 先在自己函数中查找局部变量
	// 2. 找不到就往外找全局变量
	fmt.Printf("变量num: %v\n", num)
}

func main() {
	// 函数可以作为变量，作为变量的时候没有括号
	abc := testGlobal
	abc()
	fmt.Printf("Type of abc: %T\n", abc)
}
