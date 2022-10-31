package main

import "fmt"

// 流程控制 for
func main() {
	// 1. 基本写法
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	// 2. 省略初始语句，但是必须保留初始语句后面的分号
	var j = 0
	for ; j < 10; j++ {
		fmt.Println(j)
	}

	fmt.Println()

	// 3. 省略初始语句和结束语句
	var k = 10
	for k > 0 {
		fmt.Println(k)
		k--
	}

	// 4. 无限循环
	// for {
	// 	fmt.Println("hello")
	// }
}
