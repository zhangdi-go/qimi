package main

import "fmt"

// 数组

// 数组是同一种数据类型元素的集合。 在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。

func main() {
	var a [3]int
	var b [4]int
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("*********1.数组的初始化**********")
	// 1. 数组的初始化
	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Printf("cityArray: %v\n", cityArray)
	fmt.Printf("cityArray[0]: %v\n", cityArray[0])
	fmt.Printf("cityArray[1]: %v\n", cityArray[1])
	fmt.Printf("cityArray[2]: %v\n", cityArray[2])
	fmt.Printf("cityArray[3]: %v\n", cityArray[3])

	fmt.Println("*********2.编译器推导数组长度**********")
	// 2. 编译器推导数组长度
	var boolArray = [...]bool{true, false, false, true, false}
	fmt.Printf("boolArray: %v\n", boolArray)

	fmt.Println("*********3.使用索引值方式初始化**********")
	// 3. 使用索引值方式初始化
	// 这里Go是数组的第2个元素，因为它索引值为1。前面会有一个空的字符串元素
	var langArray = [...]string{1: "Go", 2: "Ruby", 3: "Python", 6: "Java"}
	fmt.Printf("langArray: %v\n", langArray)
	// 上面的数组长度为7，类型为“[7]string”
	fmt.Printf("langArray Type: %T\n", langArray)

	fmt.Println("*******for遍历********")
	// 普通遍历
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}

	fmt.Println("*******for range遍历********")
	// for range遍历
	for _, value := range cityArray {
		fmt.Println(value)
	}
}
