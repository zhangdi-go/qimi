package main

import "fmt"

// 字符
func main() {
	// byte	uint8的别名 ASCII字符
	// rune int32的别名 UTF-8字符
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T  c2:%T\n", c1, c2)

	s := "Google中国"
	for _, r := range s {
		fmt.Printf("%c\t", r)
	}
}
