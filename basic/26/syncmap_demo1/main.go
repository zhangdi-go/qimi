package main

import (
	"fmt"
	"sync"
)

// 原生的map不是并发安全的
var m = make(map[int]int)

var wg sync.WaitGroup

func get(key int) int {
	return m[key]
}

func set(key int, value int) {
	m[key] = value
}

func main() {
	for i := 0; i < 20; i++ { // 开启20个狗揉停，并发地设置值&获取值
		wg.Add(1)
		go func(i int) {
			set(i, i+100)                              // 设置键值对
			fmt.Printf("key:%v value:%v\n", i, get(i)) // 打印键值对
			wg.Done()
		}(i)
	}
	wg.Wait()
}
