package main

import "fmt"

// 切片删除元素
func main() {
	a := []string{"北京", "上海", "广州", "深圳"}
	// 删除"广州"，索引为2，故先切片[0:2]，注意是左闭右开区间
	a = append(a[0:2], a[3:]...)
	fmt.Println(a)

	// 从切片a中删除索引为index的元素：
	// append(a[:index], a[index+1:]...)
}
