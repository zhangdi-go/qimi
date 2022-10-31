package main

import (
	"fmt"
	"sync"
)

// goroutine demo 2

var wg sync.WaitGroup // 全局定义一个wg

func main() {
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() { // 这个匿名函数还是个闭包(包含了外部的变量i)
			fmt.Println("hello,", i)
			wg.Done()
		}()
	}
	fmt.Println("hello, main")
	wg.Wait() // 阻塞, 等所有小弟都干完活再收兵
}

/*
好多个goroutine打印了相同的值
闭包要去外面找变量i，但for可能已经走了很多步了
*/
