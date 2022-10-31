package main

import (
	"fmt"
	"sync"
)

// 多个goroutine并发操作全局变量

var (
	x    int        // 全局变量x
	lock sync.Mutex // 互斥锁
	wg   sync.WaitGroup
)

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock() // 加锁
		x++
		lock.Unlock() // 释放锁
	}
	wg.Done()
}

func main() {
	wg.Add(3)
	go add()
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
