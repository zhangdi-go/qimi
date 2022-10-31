package main

import (
	"fmt"
	"sync"
)

// sync.Map 并发安全的map
var m = sync.Map{}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 20; i++ { // 开启20个狗揉停，并发地设置值获取值
		wg.Add(1)
		go func(i int) {
			m.Store(i, i+100) // 设置键值对
			value, _ := m.Load(i)
			fmt.Printf("key:%v value:%v\n", i, value) // 打印键值对
			wg.Done()
		}(i)
	}
	wg.Wait()
}
