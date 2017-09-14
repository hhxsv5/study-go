package main

import (
	"fmt"
	"runtime"
)

//Go 开发者不需要写代码来释放程序中不再使用的变量和结构占用的内存，在 Go 运行时中有一个独立的进程，即垃圾收集器（GC），会处理这些事情，它搜索不再使用的变量然后释放它们的内存。
//可以通过 runtime 包访问 GC 进程。
//通过调用 runtime.GC() 函数可以显式的触发 GC，但这只在某些罕见的场景下才有用
//比如当内存资源不足时调用 runtime.GC()，它会在此函数执行的点上立即释放一大片内存，此时程序可能会有短时的性能下降（因为 GC 进程在执行）。
func main() {
	var a int = 10
	fmt.Println(a)
	runtime.SetFinalizer(&a, func(pa *int) { //如果需要在一个对象 obj 被从内存移除前执行一些特殊操作，比如写到日志文件中，可以通过如下方式调用函数来实现：
		fmt.Println("GC a", *pa)
	})

	runtime.GC() //手动触发GC

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("LastGC：%d，已经被配并仍在使用：%dKB，从开始运行到现在分配的内存总数：%dKB，堆当前的用量：%dKB\n", m.LastGC, m.Alloc/1024, m.TotalAlloc/1023, m.HeapAlloc/1024)
}
