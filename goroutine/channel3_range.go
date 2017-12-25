package main

import "fmt"

//range和close
func xrange(start int, max int, step int, c chan int) {
	for ; start < max; start += step {
		c <- start
	}
	close(c) //显式关闭channel，记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic
	//另外channel不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的
}

func main() {
	c := make(chan int)
	go xrange(0, 10, 1, c)
	for v := range c { //for v := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭
		fmt.Println(v)
	}
}
