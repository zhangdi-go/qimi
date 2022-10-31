package main

import "fmt"

// 二维数组
func main() {
	cityArray := [4][2]string{ // cityArray有4个元素，每个元素都是[2]string这个一维数组
		{"北京", "天津"},
		{"上海", "苏州"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Printf("cityArray: %v\n", cityArray)

	fmt.Println("*******for range遍历二维数组********")
	// 先遍历外层数组
	for _, v1 := range cityArray {
		// 再遍历内层数组，取出元素
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
