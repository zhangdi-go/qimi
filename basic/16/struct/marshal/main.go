package main

import (
	"encoding/json"
	"fmt"
)

// 序列化

// 如果一个包中定义的标识符是首字母大写的，那么就是对包外可见的
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
	c1 := class{
		Title:    "火箭101",
		Students: make([]student, 0, 20), // len=0, cap=20
	}
	// 往班级c1中添加学生
	for i := 1; i <= 10; i++ {
		// 造10个学生
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}
	fmt.Printf("c1: %v\n", c1)

	fmt.Println()

	// 序列化: Go语言中的数据 -> JSON格式的字符串
	jsonStr, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
		return
	}
	fmt.Printf("jsonStr: %s\n", jsonStr)
}
