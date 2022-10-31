package main

import (
	"encoding/json"
	"fmt"
)

// 反序列化

// 如果一个Go语言包中定义的标识符是首字母大写的,那么就是对外可见的
type student struct {
	ID   int
	Name string
}

// student的构造函数
func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

// class(班级)的结构体
type class struct {
	Title    string    `json:"class_title"`
	Students []student `json:"student_list"`
}

func main() {
	var c2 class
	jsonStr := `{"class_title":"火箭101","student_list":[{"ID":1,"Name":"stu01"},{"ID":2,"Name":"stu02"},{"ID":3,"Name":"stu03"},{"ID":4,"Name":"stu04"},{"ID":5,"Name":"stu05"},{"ID":6,"Name":"stu06"},{"ID":7,"Name":"stu07"},{"ID":8,"Name":"stu08"},{"ID":9,"Name":"stu09"},{"ID":10,"Name":"stu10"}]}`
	fmt.Printf("jsonStr: %s\n", jsonStr)

	fmt.Println()

	// 反序列化: JSON格式的字符串 -> Go语言中的数据
	err := json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("json unmarshal failed, err:", err)
		return
	}
	fmt.Printf("c2: %v\n", c2)
}
