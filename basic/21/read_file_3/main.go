package main

import (
	"fmt"
	"io/ioutil"
)

// read by ioutil
func read3() {
	content, err := ioutil.ReadFile("../test3.txt") // 相对路径，(例如go run就是相对main.go的位置。编译就是相对可执行文件的位置。)
	if err != nil {
		fmt.Print("read file failed, err: ", err)
		return
	}
	fmt.Print(string(content)) // ioutil.ReadFile返回的是[]byte, 所以要转成string
}

func main() {
	read3()
}

/*
这里不用close
因为ioutil.ReadFile()调用了os.ReadFile()，在这os.ReadFile()里面已经defer了一个close
*/
