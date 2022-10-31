package main

import "fmt"

// 定义具有多个返回值的函数
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	sum, sub := calc(100, 200)
	fmt.Printf("sum: %v\n", sum)
	fmt.Printf("sub: %v\n", sub)
}
