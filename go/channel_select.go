package main

import (
	"fmt"
	"math/rand"
	"time"
)

//多个channel的时候，通过select可以监听channel上的数据流动。
//select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个channel都准备好的时候， select是随机的选择一个执行的。
//在select里面还有default语法，select其实就是类似switch的功能，当监听的channel都没有准备好的时候，默认执行的default（select不再阻塞等待channel）。

func write(c1, c2 chan int) {
	for {
		rand.Seed(time.Now().UnixNano())
		//模拟io等待
		switch rand.Int31n(3) {
		case 1:
			c1 <- int(rand.Int31n(100))
		case 2:
			c2 <- int(rand.Int31n(100))
		default:
			time.Sleep(time.Second)
		}
	}
}

func read(c1, c2 chan int) {
	a, b := 0, 0
	for {
		select {
		case a = <-c1:
			fmt.Println("read c1:", a)
		case b = <-c2:
			fmt.Println("read c2:", b)
		default:
			//fmt.Println("no ready channel")
		}
	}
}

func main() {
	c1, c2 := make(chan int), make(chan int)
	go write(c1, c2)
	read(c1, c2)
}
