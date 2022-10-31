package main

import (
	"bufio"
	"fmt"
	"os"
)

// write by bufio
func write2() {
	// 1. 打开文件
	fileObj, err := os.OpenFile("../test2.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Print("open file failed, err: ", err)
		return
	}

	// 2. 提前关闭文件
	defer fileObj.Close()

	// 3. 写入文件
	writer := bufio.NewWriter(fileObj)
	for i := 0; i < 3; i++ {
		writer.WriteString("Australia\t") // 将数据先写入缓存
	}
	writer.Flush() // 将缓存中的内容写入文件
}

func main() {
	write2()
}
