package main

import "fmt"

/*
	空接口是指没有定义任何方法的接口。
	因此任何类型都实现了空接口。
	空接口类型的变量可以存储任意类型的变量。
*/

func main() {
	// 1. 空接口类型作为函数的参数
	var x interface{}
	s := "hello, 小草莓"
	x = s
	fmt.Printf("Type:%T\tValue:%v\n", x, x)

	i := 100
	x = i
	fmt.Printf("Type:%T\tValue:%v\n", x, x)

	b := false
	x = b
	fmt.Printf("Type:%T\tValue:%v\n", x, x)

	// 2. 空接口类型可以作为map的value
	var m = make(map[string]interface{}, 8)
	m["name"] = "卓卓脑婆"
	m["age"] = 18
	m["hobby"] = []string{"唱", "跳", "rap", "篮球"}
	fmt.Printf("m: %v\n", m)
}
