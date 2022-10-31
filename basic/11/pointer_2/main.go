package main

import "fmt"

// Go的函数传值：值拷贝(ctrl+c、ctrl+v)

func modify1(x int) {
	x = 100
}

func modify2(y *int) {
	*y = 100
}

func main() {
	a := 1
	modify1(a)
	fmt.Println("a =", a) // 1
	modify2(&a)
	fmt.Println("a =", a) // 100
}
