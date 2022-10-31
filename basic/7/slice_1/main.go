package main

import "fmt"

// 切片(slice)
func main() {
	// var a []string
	// var b []int
	// var c = []bool{false, true}
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// 基于数组得到切片
	a := [5]int{55, 56, 57, 58, 59}
	b := a[1:4] // 左闭右开
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b的类型是%T\n", b)

	// 切片再次切片
	c := b[0:len(b)]
	fmt.Printf("c: %v\n", c)
	fmt.Printf("c的类型是%T\n", c)

	// make函数构造切片
	d := make([]int, 5, 10) // make([]type, len, cap)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("d的类型是%T\n", d)

	// 通过len()函数来获取切片的长度
	fmt.Print("d切片的长度是：")
	fmt.Println(len(d))
	// 通过cap()函数来获取切片的容量
	fmt.Print("d切片的容量是：")
	fmt.Println(cap(d))

	// nil
	var e []int         // 声明int型切片
	var f = []int{}     // 声明int型切片并且初始化（因为是赋值运算且后面加了个花括号）
	g := make([]int, 0) // 声明int型切片并且初始化（注意海象运算符）
	// e、f、g在下面三行打印出来都是[]，len是0，cap是0  但是只有e是nil
	if e == nil {
		fmt.Println("e==nil")
	}
	fmt.Println("e:", e, len(e), cap(e))
	if f == nil {
		fmt.Println("f==nil")
	}
	fmt.Println("f:", f, len(f), cap(f))
	if g == nil {
		fmt.Println("g==nil")
	}
	fmt.Println("g:", g, len(g), cap(g))
	fmt.Println()

	// ***切片的赋值拷贝***
	i := make([]int, 3) // i=[0 0 0]
	j := i              // i和j共用一个(底层数组)地址
	j[0] = 100          // 修改下标为0的元素的值，俩全都改变
	fmt.Printf("i: %v\n", i)
	fmt.Printf("j: %v\n", j)
	fmt.Println()

	// 切片的遍历
	k := []int{1, 2, 3, 4, 5}

	// 1. 根据索引来遍历
	for i := 0; i < len(k); i++ {
		fmt.Println(i, k[i])
	}
	fmt.Println()
	// 2. 使用for range遍历
	for i, v := range k {
		fmt.Println(i, v)
	}
}
