package test

import "testing"

//go test
//go test -file xx.go
//go test -run=Testxxx
//go test -v 不加"-v"表示只显示未通过的用例结果

func TestSum(t *testing.T) {
	result := Sum(-101, 100)
	t.Log(result)
}

func TestDivide(t *testing.T) {
	result := Divide(1, 100)
	t.Log(result)
}

//go test -bench=".*"
//go test -bench=".*" -cpuprofile=cpu.out 表示生成的cpu profile文件
//go tool pprof test.test cpu.out 使用go tool pprof工具分析cpu.out
//help查看更多命令
//png 生成一张调用图的png图
//web生成一个分析的html
//text 调用图
//tree 树形调用图

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(int32(i), int32(i+100))
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(float32(i), float32(i+1000))
	}
}
