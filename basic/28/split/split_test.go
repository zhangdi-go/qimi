package split

import (
	"reflect"
	"testing"
)

/*
所有以“_test.go”结尾的源代码文件都是go test的一部分，不会被go build编译到最终的可执行文件中去。
go test命令会遍历所有的_test.go文件中符合规范的函数，然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。
*/

func TestSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string // 字符串按照sep分割
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试，以便在go test -v中看RUN的详情
			got := Split(testcase.input, testcase.sep)
			// 切片不能直接比较，所以调用反射包里面的DeepEqual
			if !reflect.DeepEqual(got, testcase.want) {
				// 测试用例失败了
				t.Errorf("name:%s failed, want:%v, got:%v", name, testcase.want, got)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	// b.N 不是固定的数
	for i := 0; i < b.N; i++ {
		Split("我爱你", "爱")
	}
}
