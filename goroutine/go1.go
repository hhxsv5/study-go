package main

import (
	"io/ioutil"
	"os"
	"fmt"
)

var exit = make(chan struct{}) //创建通道，只能由make创建

func logfile(s string) {
	fmt.Println(s)
	ioutil.WriteFile("/Users/dave/Documents/www/go/src/test/go/go1.log", []byte(s), os.ModePerm)
	close(exit) //关闭通道
}

func main() {
	//使用go关键字：创建一个并发任务单元
	//新建任务被放置在系统队列中，等待调度器安排合适系统线程去获取执行权。
	//当前流程不会阻塞，不会等待该任务启动，且运行时也不保证并发任务的执行次序。
	//与defer一样，goroutine也会因“延迟执行”而立即计算并复制执行参数。
	//go ioutil.WriteFile("/Users/dave/Documents/www/go/src/test/go/go1.log", []byte("go write file"), os.FileMode(os.O_CREATE|os.O_WRONLY))

	go logfile("go func write file")

	//进程退出时不会等待并发任务结束，可用通道(channel)阻塞，然后发出退出信号。
	//time.Sleep(time.Second * 3)
	fmt.Println("start")

	<-exit //通道未关闭前一直阻塞，直到关闭，则结束阻塞

	fmt.Println("exit")
}
