package main

import "fmt"

// 运算符
func main() {
	// 1. 算术运算符
	a := 5
	b := 2
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // 2
	fmt.Println(a % b) // 1

	fmt.Println("***********************")

	// 2. 关系运算符
	fmt.Println(10 > 2)  // true
	fmt.Println(10 != 2) // true
	fmt.Println(4 <= 5)  // true

	fmt.Println("***********************")

	// 3. 逻辑运算符
	fmt.Println(10 > 5 && 1 == 1) // true
	fmt.Println(!(10 > 5))        // false
	fmt.Println(1 > 5 || 1 == 1)  // true

	fmt.Println("***********************")

	// 4. 位运算符
	c := 1             // 001
	d := 5             // 101
	fmt.Println(c & d) // 001 => 1 按位与
	fmt.Println(c | d) // 101 => 5 按位或
	fmt.Println(c ^ d) // 100 => 4 按位异或

	fmt.Println("***********************")

	// 5. 赋值运算符
}
