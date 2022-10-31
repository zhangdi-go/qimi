package main

import "fmt"

// 函数外的每个语句都必须以关键字开始（var、const、func等）

// 变量

// 1.1 标准声明格式
var name string
var age int
var isOk bool

// 1.2 批量声明变量
var (
	a string
	b int
	c bool
	d float32
)

// 2. 声明变量的同时指定初始值
var name1 string = "小明"
var age1 int = 18

// 3. 类型推导，让编译器根据变量的初始值推导出其类型
var name2, age2 = "小红", 19

func main() {
	fmt.Println(name, age, isOk)
	fmt.Println(a, b, c, d)
	fmt.Println(name1, age1)
	fmt.Println(name2, age2)

	// 4. 短变量声明，只能用在函数中
	m := 10
	fmt.Println(m)
}
