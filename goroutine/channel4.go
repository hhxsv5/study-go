package main

func main() {
	c := make(chan int, 3)

	//将 channel 隐式转换为单向队列，只收或只发。
	send := (chan<- int)(c) //只能发，等同于var send chan<- int = c
	recv := (<-chan int)(c) //只能收，等同于var recv <-chan int = c

	send <- 1
	//<-send//报错

	<-recv
	//recv<-1//报错
}
