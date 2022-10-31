package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁(适合读多写少的场景)

var (
	x      int          // 全局变量x
	lock   sync.Mutex   // 互斥锁
	rwLock sync.RWMutex // 读写锁
	wg     sync.WaitGroup
)

func read() {
	wg.Add(1)
	//lock.Lock() // 加互斥锁
	rwLock.RLock() // 加读锁
	time.Sleep(time.Millisecond)
	//lock.Unlock() // 释放互斥锁
	rwLock.RUnlock() // 释放读锁
	wg.Done()
}

func write() {
	wg.Add(1)
	//lock.Lock() // 加互斥锁
	rwLock.Lock() // 加写锁
	x++
	time.Sleep(time.Millisecond * 10)
	//lock.Unlock() // 释放互斥锁
	rwLock.Unlock() // 释放写锁
	wg.Done()
}

func main() {
	start := time.Now()

	for i := 0; i < 100; i++ {
		go read()
	}

	for i := 0; i < 10; i++ {
		go write()
	}

	wg.Wait()

	fmt.Println(time.Now().Sub(start))
}
