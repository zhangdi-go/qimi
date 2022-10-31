package main

import (
	"fmt"
	"reflect"
)

// set value

func reflectSetValue(x any) {
	v := reflect.ValueOf(x)
	switch v.Elem().Kind() { // Elem()用来获取指针取对应的值
	case reflect.Int32:
		v.Elem().SetInt(123)
	case reflect.Float32:
		v.Elem().SetFloat(123.456)
	}
}

func main() {
	var aaa int32 = 0
	reflectSetValue(&aaa) // 传的是指针
	fmt.Println(aaa)
}
