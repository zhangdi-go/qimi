package main

import "fmt"

// 数组是值类型，传入的参数只是一个副本
func modify1(a [3]int) {
	a[0] = 100 // 数组副本的a[0]被修改为100
}

// 传入数组的指针
func modify2(a *[3]int) {
	a[0] = 100
}

func main() {
	x := [3]int{1, 2, 3}
	fmt.Printf("x的原始值: %v\n", x)
	modify1(x) // 修改不会成功
	fmt.Printf("x经过modify1修改后: %v\n", x)
	modify2(&x) // 修改会成功
	fmt.Printf("x经过modify2修改后: %v\n", x)
}
