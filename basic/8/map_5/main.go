package main

import "fmt"

// 元素类型为map的slice

func main() {
	var mapSlice = make([]map[string]int, 8, 8) // 只完成了切片的初始化
	// [nil nil nil nil nil nil nil nil]
	fmt.Printf("mapSlice: %v\n", mapSlice)
	fmt.Println(mapSlice[0] == nil) // true

	mapSlice[0] = make(map[string]int, 8) // 完成了map元素的初始化
	mapSlice[0]["小米"] = 3999
	fmt.Printf("mapSlice: %v\n", mapSlice)
}
