package main

import "fmt"

// 流程控制 switch
func main() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入")
	}

	// case中使用表达式
	age := 30
	switch {
	case age > 18:
		fmt.Println("澳门首家线上赌场开业了")
	case age < 18:
		fmt.Println("warning")
	default:
		fmt.Println("嘿嘿嘿")
	}
}
