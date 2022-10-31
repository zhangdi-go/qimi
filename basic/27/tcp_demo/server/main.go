package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp server demo

func process(conn net.Conn) {
	// 处理完之后要关闭这个连接
	defer conn.Close()
	// 针对当前的连接做数据的发送和接收操作
	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Print("read from conn failed, err: ", err)
			// 结果出错就跳出for循环
			break
		}
		receive := string(buf[:n])
		fmt.Println("接收到的数据:", receive)
		conn.Write([]byte("ok")) // 把收到的数据返回给客户端
	}
}

func main() {
	// 1. 开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Print("listen failed, err: ", err)
		return
	}
	for {
		// 2. 等待客户端来建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Print("accept failed, err: ", err)
			// 让for走下一次循环，而不是跳出for循环
			continue
		}
		// 3. 启动一个单独的go routine去处理连接
		go process(conn)
	}
}
