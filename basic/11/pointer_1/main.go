package main

import "fmt"

/*
1. &(取地址)、*(根据地址取值)
2. 区别于C/C++中的指针，Go语言中的指针不能进行偏移和运算，是安全指针。
3. Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：*int、*int64、*string等。
*/

func main() {
	a := 10
	b := &a
	fmt.Printf("%v %p\n", a, &a)
	fmt.Printf("%v %p\n", *b, b)
	// 变量b是一个int类型的指针(*int)，它存储的是变量a的内存地址
	c := &b                                    // 取变量b的地址
	fmt.Printf("Type c:%T Value c:%v\n", c, c) // Type c:**int
}
