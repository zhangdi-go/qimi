package main

import (
	"fmt"
	"io"
	"os"
)

// read by os.Read()
func read1() {
	// 1. 打开文件
	fileObj, err := os.Open("../test1.txt") // 相对路径，(例如go run就是相对main.go的位置。编译就是相对可执行文件的位置。)
	if err != nil {
		fmt.Print("open file failed, err: ", err)
		return
	}

	// 2. 提前关闭文件
	defer fileObj.Close()

	// 3. 使用os.Read()读取文件
	for {
		var tmp = make([]byte, 1)
		n, err := fileObj.Read(tmp)
		if err == io.EOF { // End Of File 表示读到了文件的末尾
			return
		}
		if err != nil {
			fmt.Print("read file failed, err: ", err)
			return
		}
		fmt.Print(string(tmp[:n]))
	}
}

func main() {
	read1()
}
