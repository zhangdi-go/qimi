package main

import (
	"fmt"
	"strings"
)

// 字符串常见操作
func main() {
	// 求字符串的长度，len()方法
	s1 := "hello"
	fmt.Println(len(s1))
	// 汉字，每个字占3字节
	s2 := "中国"
	fmt.Println(len(s2))

	// 字符串拼接
	fmt.Println(s1 + s2)
	s3 := fmt.Sprintf("%s%s", s1, s2)
	fmt.Println(s3)
	// 字符串分割
	s4 := "how do you do"
	fmt.Println(strings.Split(s4, " "))
	// join
	s5 := []string{"how", "do", "you", "do"}
	fmt.Println(s5)
	fmt.Println(strings.Join(s5, "+"))
}
