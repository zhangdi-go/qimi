package main

import "fmt"

func main() {
	// 判断某个键存不存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["小明"] = 100
	scoreMap["小红"] = 200

	// 判断"李华"在不在scoreMap中
	v, ok := scoreMap["小强"] // 如果map中存在这个key：那么ok返回true，map的value就赋值给v；否则ok返回false，map的value的零值赋值给v。
	if ok {
		fmt.Println("查到了，他的分数是:", v)
	} else {
		fmt.Println("查无此人，他的分数是:", v)
	}
}
