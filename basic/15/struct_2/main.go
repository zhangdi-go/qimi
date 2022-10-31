package main

import "fmt"

// 结构体的嵌套

// 用type关键字造的，都是类型

type Address struct {
	Province string
	City     string
}

type Person struct {
	Name   string
	Gender string
	Age    int
	//Address Address // 嵌套另外一个结构体。前面的是Person里面的字段名，后面的是指Address类型
	Address // 匿名嵌套
}

func main() {
	p1 := Person{
		Name:   "小草莓",
		Gender: "男",
		Age:    18,
		Address: Address{
			Province: "上海",
			City:     "浦东",
		},
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.Name, p1.Gender, p1.Age, p1.Address)
	fmt.Println(p1.Address.City) // 通过嵌套的结构体访问其内部的字段
	fmt.Println(p1.City)         // 直接访问匿名结构体中的字段
}
