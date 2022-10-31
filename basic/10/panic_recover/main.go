package main

import "fmt"

// panic和recover

// 注意：
// 1. recover()必须搭配defer使用。
// 2. defer一定要在可能引发panic的语句之前定义。

func a() {
	fmt.Println("func a")
}

func b() {
	defer func() {
		err := recover() // defer里用一个recover接收错误
		if err != nil {
			fmt.Printf("func b error, the err is: \"%v\"\n", err)
		}
	}()
	panic("panic in b")
}

func c() {
	fmt.Println("func c")
}

func main() {
	a()
	b()
	c()
}
