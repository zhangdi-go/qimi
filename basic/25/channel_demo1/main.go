package main

import "fmt"

func main() {
	var ch1 chan int        // 引用类型, 需要初始化之后使用
	ch1 = make(chan int, 1) // 无缓冲区通道==同步通道，有缓冲区通道==异步通道
	ch1 <- 10               // 箭头的指向就是数据的流向
	x := <-ch1              // 箭头的指向就是数据的流向
	fmt.Println(x)
	close(ch1) // 关闭通道（不是必须的，通道是可以被GC回收的）
}
