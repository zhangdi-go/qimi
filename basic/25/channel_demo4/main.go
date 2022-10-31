package main

import "fmt"

/*
select多路复用
可处理一个或多个channel的发送/接收操作。
如果多个case同时满足，select会随机选择一个。
对于没有case的select会一直等待，可用于阻塞main函数。
*/

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Println("啥都不干")
		}
	}
}
