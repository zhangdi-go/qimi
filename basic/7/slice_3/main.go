package main

import "fmt"

// 切片复制
func main() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5, 5)
	c := b     // b直接赋值给c，b和c共用一个底层数组
	copy(b, a) // 把a复制给b
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Println() // 空一行
	// 复制了，a和b已经是两个不同的了
	b[0] = 100
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
}
