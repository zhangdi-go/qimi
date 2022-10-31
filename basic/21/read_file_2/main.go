package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// read by bufio
func read2() {
	// 1. 打开文件
	fileObj, err := os.Open("../test2.txt") // 相对路径，(例如go run就是相对main.go的位置。编译就是相对可执行文件的位置。)
	if err != nil {
		fmt.Print("open file failed, err: ", err)
		return
	}

	// 2. 提前关闭文件
	defer fileObj.Close()

	// 3. 使用bufio来读取文件
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print(line)
			return
		}
		if err != nil {
			fmt.Print("read file failed, err: ", err)
			return
		}
		fmt.Print(line)
	}
}

func main() {
	read2()
}
