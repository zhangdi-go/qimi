package main

import "fmt"

// 切片扩容
func main() {
	// 切片要初始化后才能使用
	// append()函数将元素追加到数组的最后并返回该数组
	var a []int // 此时它并没有申请内存
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", a, len(a), cap(a), a) // 切片是引用类型，所以最后面的不用写取地址符号
	}
	fmt.Println()

	// append还可以追加多个元素
	var b []int // 此时它并没有申请内存
	b = append(b, 1, 2, 3, 4, 5)
	fmt.Println("b:", b)
	c := []int{12, 13, 14, 15}
	b = append(b, c...)
	fmt.Println("b:", b)
}
