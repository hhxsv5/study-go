package main

import (
	"fmt"
	"runtime"
)

func main() {
	//退出当前执行的goroutine，但是defer函数还会继续调用。
	//该函数不影响其他并发任务，不会引发panic，无法捕获。
	exit := make(chan struct{})

	go func() {
		defer close(exit)      //执行
		defer fmt.Println("a") //执行

		func() {
			defer func() {
				fmt.Println("b", recover() == nil) //执行，recover返回nil
			}()

			func() {
				defer fmt.Println("c") //执行
				runtime.Goexit()       //终止当前协程任务，调用defer
				fmt.Println("c done")  //不会执行
			}()

			fmt.Println("b done") //不会执行
		}()

		fmt.Println("a done") //不会执行
	}()

	<-exit

	fmt.Println("main exit")

	/**
		c
		b true
		a
		main exit

	 */

}
