package main

import "fmt"

// 结构体本质上是一种聚合型的数据类型
type person struct {
	name, city string // 行尾不用写逗号
	age        int    // 行尾不用写逗号
}

func main() {
	var p person
	p.name = "小草莓"
	p.age = 26
	p.city = "上海"
	fmt.Printf("姓名:%s\n年龄:%d\n城市:%s\n", p.name, p.age, p.city)
	fmt.Printf("p: %v\n", p)
}
