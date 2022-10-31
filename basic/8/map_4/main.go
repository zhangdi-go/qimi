package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 按照某个固定顺序遍历map(思路：取出key存储到一个切片中，然后对key排序并使用其对map检索)

func main() {
	var scoreMap = make(map[string]int, 100)

	// 添加50个键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100) // 0~99的随机整数
		scoreMap[key] = value
	}

	// 按照key从小到大的顺序去遍历scoreMap
	// 1. 先取出所有的key存放到切片keys中
	keys := make([]string, 0, 100) // 注意此处make()里第二个参数是切片的len，一定设置成0，否则会制造出空字符串
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// 2. 对keys做排序
	sort.Strings(keys) // keys目前是一个有序的切片
	// 3. 按照排序后的keys对scoreMap检索
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
	// but，为什么随机生成了固定的数？甚至和视频里一样？——————>因为没有使用随机数种子
}
