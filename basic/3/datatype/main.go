package main

import (
	"fmt"
	"math"
)

// 基本数据类型
func main() {
	// 十进制转二进制
	n := 10
	fmt.Printf("%b\n", n) // 1010
	fmt.Printf("%d\n", n) // 10

	// 八进制
	m := 075
	fmt.Printf("%d\n", m) // 61
	fmt.Printf("%o\n", m) // 75

	// 十六进制
	f := 0xff
	fmt.Printf("%d\n", f) // 255
	fmt.Printf("%x\n", f) // ff

	fmt.Println("********分割线*********")

	// 浮点型：go只有float32和float64，没有double！
	// IEEE 754标准
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	fmt.Println("********分割线*********")

	// 布尔型
	var a bool
	fmt.Println(a) // false
	a = true
	fmt.Println(a) // true

	fmt.Println("********分割线*********")

	// 字符串
	// 打印Windows下的一个路径“C:\go\bin\go.exe”
	fmt.Println("C:\\go\\bin\\go.exe")
	// 多行字符串，使用反引号``来包裹（反引号需要在英文模式下输入）
	s1 := `
第一行
	第二行
		第三行
			原样输出
`
	fmt.Println(s1)
}
