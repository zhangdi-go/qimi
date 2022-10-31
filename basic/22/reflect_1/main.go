package main

import (
	"fmt"
	"reflect"
)

// type of

func reflectType(x any) {
	// 我是不知道别人调用我这个函数的时候会传进来什么类型的变量的
	// 方式1：通过类型断言,一个个猜
	// 方法2：反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Kind())
}

type Cat struct {
	name string
}

func main() {
	var a int8 = 1
	reflectType(a)
	var b float32 = 1.234
	reflectType(b)
	var c Cat
	reflectType(c)
	var d []int
	reflectType(d)
	var e []string
	reflectType(e)
}
