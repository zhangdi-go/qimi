package main

import "fmt"

func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}               // s = [1 2]
	s = append(s, 3)               // s = [1 2 3]
	fmt.Printf("%+v\n", s)         // s = [1 2 3]
	m["q1mi"] = s                  // m["q1mi"] = [1 2 3]
	s = append(s[:1], s[2:]...)    // 删除掉s中索引为1的元素'2'
	fmt.Printf("%+v\n", s)         // s = [1 3]
	fmt.Printf("%+v\n", m["q1mi"]) // m["q1mi"] = [1 3 3]
}
