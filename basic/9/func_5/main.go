package main

import "fmt"

// 函数可以作为参数

func sum(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	r1 := calc(100, 200, sum)
	fmt.Printf("r1: %v\n", r1)
	r2 := calc(100, 200, sub)
	fmt.Printf("r2: %v\n", r2)
}
