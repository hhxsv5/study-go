package main

import "fmt"

//golang的channel实现生产消费者模型
func producer(c chan int, max int) {
	defer close(c)

	for i := 0; i < max; i++ {
		c <- i
	}
}

func consumer(c chan int) {
	var d int
	var ok = true
	for ok {
		if d, ok = <-c; ok {
			fmt.Println(d)
		}
	}
}

func main() {
	c := make(chan int)
	go producer(c, 10) //将producer函数交给sub goroutine处理, 产生的结果传入channel中
	consumer(c)        //main goroutine从channel中取数据
}
