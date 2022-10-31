package main

import (
	"fmt"
	"reflect"
)

// value of

func reflectValue(x any) {
	v := reflect.ValueOf(x)
	// 如何得到一个传入时候类型的变量
	switch v.Kind() { // 拿到值对应的类型的种类
	case reflect.Int32:
		// 把反射取到的值转换成一个int32类型的变量
		ret := int32(v.Int())
		fmt.Printf("Type: %T Value: %v\n", ret, ret)
	case reflect.Float32:
		// 把反射取到的值转换成一个float32类型的变量
		ret := float32(v.Float())
		fmt.Printf("Type: %T Value: %v\n", ret, ret)
	}
}

func main() {
	var aa int32 = 1
	reflectValue(aa)
	var bb float32 = 1
	reflectValue(bb)
}
