package main

import "fmt"

// 自定义类型和类型别名

// MyInt 基于int类型的自定义类型，自定义类型是定义了一个全新的类型。
type MyInt int

// NewInt int类型的别名，本质上NewInt与int是同一个类型。
type NewInt = int

func main() {
	var a MyInt
	var b NewInt
	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of b: %T\n", b)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
}
