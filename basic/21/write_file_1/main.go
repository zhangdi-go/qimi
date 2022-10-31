package main

import (
	"fmt"
	"os"
)

// write by os.Write()
func write1() {
	// 1. 打开文件
	fileObj, err := os.OpenFile("../test1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Print("open file failed, err: ", err)
		return
	}

	// 2. 提前关闭文件
	defer fileObj.Close()

	// 3. 写入文件
	str := "一二三四五\n"
	fileObj.Write([]byte(str))     // []byte
	fileObj.WriteString("六七八九十\n") // string
}

func main() {
	write1()
}
