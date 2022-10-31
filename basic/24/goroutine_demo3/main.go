package main

import (
	"fmt"
	"sync"
)

// goroutine demo 3

var wg sync.WaitGroup // 全局定义一个wg

func main() {
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func(i int) { // 这个匿名函数还是个闭包(包含了外部的变量i)
			fmt.Println("hello,", i)
			wg.Done()
		}(i) // 显式地传一个副本进来, 固定下来, 保证一致性
	}
	fmt.Println("hello, main")
	wg.Wait() // 阻塞, 等所有小弟都干完活再收兵
}
