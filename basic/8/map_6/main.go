package main

import "fmt"

// 值为slice的map

func main() {
	var sliceMap = make(map[string][]int, 8) // 只完成了map的初始化
	v, ok := sliceMap["中国"]
	if ok {
		fmt.Println(v)
	} else {
		// sliceMap中没有中国这个键，进行对值的切片初始化
		sliceMap["中国"] = make([]int, 8, 8) // 完成了对slice的初始化
		sliceMap["中国"][1] = 100
		sliceMap["中国"][2] = 200
		sliceMap["中国"][3] = 300
	}
	// 遍历sliceMap
	fmt.Println("开始遍历sliceMap......")
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}
}
