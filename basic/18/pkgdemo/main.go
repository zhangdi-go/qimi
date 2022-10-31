package main

import (
	"fmt"
	"pkgdemo/calc"
)

func main() {
	fmt.Println("hello", calc.Name)
	fmt.Println(calc.Add(100, 200))
}

func init() {
	fmt.Println("main.init()")
}
