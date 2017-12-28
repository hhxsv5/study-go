package main

import (
	"fmt"
	"sync"
	"io/ioutil"
	"runtime"
)

func readIO() string {
	c, e := ioutil.ReadFile("./io.data")
	if e != nil {
		panic(e)
	}
	return string(c)
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1) //等待1个goroutine完成

	go func() {
		defer wg.Done()
		defer fmt.Println("run defer") //Goexit()后会执行

		data := readIO()
		fmt.Println("io", data)

		runtime.Goexit() //终止当前goroutine

		fmt.Println("after Goexit()") //不会执行
	}()

	wg.Wait()

	fmt.Println("goroutine done")
}
