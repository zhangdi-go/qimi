package main

import "fmt"

/*
iota是go语言的常量计数器，只能在常量的表达式中使用。
iota在const关键字出现时将被重置为0。
const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
使用iota能简化定义，在定义枚举时很有用。
*/
const (
	n1 = iota // 0
	n2        // 1
	n3        // 2
	n4        // 3
)

const (
	n5 = iota // 0
	n6        // 1
	_         // 占位
	n8        // 3
)

const (
	n9  = iota // 0
	n10 = 100  // 100 插队
	n11 = iota // 2
	n12        // 3
)

// 常量之iota
func main() {
	fmt.Println(n1, n2, n3, n4)
	fmt.Println(n5, n6, n8)
	fmt.Println(n9, n10, n11, n12)
}
