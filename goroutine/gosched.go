package main

import (
	"fmt"
	"runtime"
)

func main() {
	//Gosched让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。
	//当前线程被放回队列，等待下次调度时恢复执行。
	//此函数很少被使用，只是在主动切换时会被使用

	//runtime.NumGoroutine()//返回正在执行和排队的任务总数
	runtime.GOMAXPROCS(1) //设置仅适用1个CPU，模拟队列任务调度
	exit := make(chan struct{})
	go func() {
		defer close(exit)
		fmt.Println("go a")

		go func() {
			fmt.Println("go b")
		}()

		for i := 0; i < 5; i++ {
			fmt.Println("go a:", i)
			if i == 2 {
				runtime.Gosched() //暂停当前任务，释放线程去执行任务b
			}
		}
	}()

	<-exit
}
