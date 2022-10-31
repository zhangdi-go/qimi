package calc

import (
	"fmt"
)

var Name = "小草莓"

func Add(x, y int) int {
	return x + y
}

// init函数在包导入的时候自动执行
// 执行顺序: 全局声明 -> init() -> main()
// init函数没有参数, 没有返回值
func init() {
	fmt.Println("add.init()")
}
