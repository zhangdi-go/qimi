package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// udp client demo

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Print("dial failed, err: ", err)
		return
	}
	defer conn.Close()
	// 从系统的标准输入读取数据
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		_, err = conn.Write([]byte(s)) // 网络操作，发送和接收的都是字节
		if err != nil {
			fmt.Print("send to server failed, err: ", err)
			return
		}
		// 接收数据
		var buf [1024]byte
		n, addr, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Print("recv from udp failed, err: ", err)
			return
		}
		fmt.Printf("read from %v, msg: %v", addr, string(buf[:n]))
	}
}
