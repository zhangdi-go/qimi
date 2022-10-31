package main

import (
	"fmt"
	"io/ioutil"
)

// write by ioutil
func write3() {
	str := "New Zealand"
	err := ioutil.WriteFile("../test3.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err: ", err)
		return
	}
}

func main() {
	write3()
}
