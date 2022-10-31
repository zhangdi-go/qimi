package main

// package main表示这是个可独立执行的程序，不是库

// 导入内置的fmt包
import "fmt"

func main() {
	fmt.Print("人生苦短，let's Go！")
}

/*
1. go语言大括号不换行
2. 语句结束不写分号
3. go install表示安装的意思，
	它先编译源代码得到可执行文件，然后将可执行文件移动到GOPATH的bin目录下。
	因为环境变量中配置了GOPATH下的bin目录，所以我们就可以在任意地方直接执行可执行文件了。
*/
