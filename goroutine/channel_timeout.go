package main

import (
	"fmt"
	"time"
)

//有时候会出现goroutine阻塞的情况，通过select来设置超时
func main() {

	//channel c是数据通道，channel o是超时判断通道
	//当sub goroutine读取c 5秒超时后，写入channel o一个true的标记，main goroutine读取到这个标记后可以做相应的处理
	c, o := make(chan int), make(chan bool)

	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(time.Second * 5):
				fmt.Println("sub goroutine read channel c timeout")
				o <- true
				break
			}
		}
	}()

	if <-o {
		//some handle
		fmt.Println("main goroutine capture channel timeout")
	}
}
