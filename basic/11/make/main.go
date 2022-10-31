package main

import "fmt"

/*
make只用于slice、map、channel的初始化，返回的还是这三个引用类型本身。函数签名如下：
	func make(t Type, size ...IntegerType) Type
make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。
*/

func main() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["小草莓"] = 100
	fmt.Printf("b: %#v\n", b)
}
