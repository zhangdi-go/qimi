package split

import "strings"

// 单元测试

// Split 将str按照sep进行分割, 返回一个字符串的切片
// Split("我爱你", "爱") => ["我", "你"]
func Split(str, sep string) (ret []string) {
	// 优化，减少内存申请次数
	// 使用make函数统一的一次初始化内存
	// count一下，str中有多少个sep那么大的空间，再加1
	ret = make([]string, 0, strings.Count(str, sep)+1)
	// idx: sep在str中的索引
	idx := strings.Index(str, sep)
	for idx >= 0 {
		// 索引前面的追加到ret里
		ret = append(ret, str[:idx])
		// 切到最后。
		// 按照sep切割，但sep不一定只占1个字节，所以这里使用len(sep)获取sep的长度
		str = str[idx+len(sep):]
		// 更新索引
		idx = strings.Index(str, sep)
	}
	// 剩下的追加到ret中
	ret = append(ret, str)
	return
}
