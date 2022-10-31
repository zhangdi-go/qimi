package main

import (
	"fmt"
	"reflect"
)

// 结构体反射

/*
反射是把双刃剑:
	1. 基于反射的代码是很脆弱的, 反射中的错误会在真正运行的时候才会引发panic
	2. 大量使用反射的代码通常难以理解
	3. 反射的性能低下
*/

type student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

func main() {
	stu1 := student{
		Name:  "小草莓",
		Score: 90,
	}

	// 通过反射去获取结构体中的所有字段信息
	t := reflect.TypeOf(stu1)
	fmt.Printf("name: %v kind: %v\n", t.Name(), t.Kind())
	fmt.Println("------------------------------")

	// 遍历结构体变量的所有字段
	for i := 0; i < t.NumField(); i++ {
		// 根据结构体字段的索引去取字段
		fieldObj := t.Field(i)
		fmt.Printf("name: %v, type: %v, tag: %v\n", fieldObj.Name, fieldObj.Type, fieldObj.Tag)
		fmt.Println("json tag is:", fieldObj.Tag.Get("json"), ", ini tag is:", fieldObj.Tag.Get("ini"))
	}
	fmt.Println("------------------------------")

	// 根据名字去取结构体中的字段
	fieldObj2, ok := t.FieldByName("Score")
	if ok {
		fmt.Println("ok, 取到了这个字段:")
		fmt.Printf("name: %v, type: %v, tag: %v\n", fieldObj2.Name, fieldObj2.Type, fieldObj2.Tag)
	}
}
