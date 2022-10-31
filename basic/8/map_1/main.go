package main

import (
	"fmt"
)

// map(映射)

func main() {
	// 1. 声明但不初始化,map的零值就是nil
	var a map[string]int
	fmt.Println(a == nil)
	// 2. map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)
	// 3. map中添加键值对
	a["小明"] = 18
	a["小强"] = 20
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a Type: %T\n", a) // a Type: map[string]int

	// 声明map的同时,完成初始化
	b := map[int]bool{
		1: true,
		2: false,
		3: true,
	}
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b Type: %T\n", b) // b Type: map[int]bool
}
