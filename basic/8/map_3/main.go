package main

import "fmt"

// 遍历map

func main() {
	// map是无序的，键值对和添加的顺序无关
	var scoreMap = make(map[string]int, 8)
	scoreMap["红米"] = 1999
	scoreMap["小米"] = 3999
	scoreMap["华为"] = 5999

	// for range
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	// 只遍历key
	for k := range scoreMap {
		fmt.Println(k)
	}
	// 只遍历value
	for _, v := range scoreMap {
		fmt.Println(v)
	}

	// delete()删除键值对
	delete(scoreMap, "华为")
	fmt.Println("删除之后:", scoreMap)
}
