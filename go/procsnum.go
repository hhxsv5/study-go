package main

import (
	"runtime"
	"fmt"
)

func main() {
	//运行时可能会创建很多线程，但任何时候仅有限的几个线程参与并发任务执行。
	//该数量默认与处理器核数相等，可使用runtime.GOMAXPROCS函数(或环境变量)修改。
	//如果参数小于1，会返回当前设置值，不做任何调整。
	//runtime.NumCPU函数可返回当前机器的核数。

	fmt.Println("current max procs", runtime.GOMAXPROCS(0))
	fmt.Println("current cpu num", runtime.NumCPU())

	runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	fmt.Println("current max procs", runtime.GOMAXPROCS(0))
}
